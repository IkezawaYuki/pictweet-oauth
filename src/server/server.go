package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/IkezawaYuki/pictweet-oauth/src/authpb"
	"github.com/IkezawaYuki/pictweet-oauth/src/infrastructure"
	server "github.com/IkezawaYuki/pictweet-oauth/src/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

var (
	authEndpoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of YourService")
)

// StartGrpcServer gRPCエントリーポイントの起動
func StartGrpcServer() {
	fmt.Println("start grpc server entry point...")
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		panic(err)
	}

	c := infrastructure.GetConnection()
	defer c.Close()

	service := server.NewAuthService(infrastructure.NewRedisHandler(c))
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, service)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}

func StartHttpServer(){
	fmt.Println("start reverse proxy server entry point...")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *authEndpoint , opts)
	if err != nil{
		log.Fatalf("failed to reverse: %v", err)
	}

	if err := http.ListenAndServe(":8080", mux); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}

}