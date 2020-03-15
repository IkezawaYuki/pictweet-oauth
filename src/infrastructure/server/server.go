package server

import (
	"context"
	"flag"
	"github.com/IkezawaYuki/pictweet-oauth/src/authpb"
	"github.com/IkezawaYuki/pictweet-oauth/src/infrastructure/datastore"
	server "github.com/IkezawaYuki/pictweet-oauth/src/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	authEndpoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of authService")
)

func Run() {
	gomniauth.SetSecurityKey(os.Getenv("SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(os.Getenv("CLIENT_ID"), os.Getenv("SECRET_VALUE"), "http://localhost:8080/auth/callback/google"),
	)

	go startGrpcEndPoint()
	startRestEndPoint()
}

// startGrpcEndPoint gRPCエントリーポイントの起動
func startGrpcEndPoint() {
	log.Println("start grpc server entry point...")
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		panic(err)
	}

	c := datastore.GetConnection()
	defer c.Close()

	service := server.NewAuthService(datastore.NewRedisHandler(c))
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, service)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}

func startRestEndPoint() {
	log.Println("start reverse proxy server entry point...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *authEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to reverse: %v", err)
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
