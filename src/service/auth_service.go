package server

import (
	"context"
	"fmt"
	"log"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"github.dip-net.co.jp/dip/robotics-auth/src/authpb"
	"github.dip-net.co.jp/dip/robotics-auth/src/infrastructure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthService is ...
type AuthService struct {
	redisHandler infrastructure.RedisHandler
}

// Login ログイン処理の実行　Googleのプロバイダを使用。
func (s *AuthService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatalln("failed to get provider: ", provider, "-", err)
	}
	loginURL, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Fatalln("error is occured when calling GetBeginAuthURL methods: ", provider, "-", err)
	}
	header := metadata.Pairs("Location", loginURL)
	grpc.SendHeader(ctx, header)

	return &authpb.LoginResponse{}, nil
}

// CallBack Googleでの認証後の処理
func (s *AuthService) CallBack(ctx context.Context, req *authpb.CallBackRequest) (*authpb.CallBackResponse, error) {
	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatalln("failed to get provider", provider, "-", err)
	}
	query := req.GetCode()
	fmt.Println(query)
	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(req.GetCode()))
	if err != nil {
		log.Fatalln("authentication could not be completed.")
	}
	user, err := provider.GetUser(creds)
	if err != nil {
		log.Fatalln("failed to get user", provider, "-", err)
	}
	// todo redisへの登録
	fmt.Println(user)

	s.redisHandler.SetWithExpire("test", "test")

	return &authpb.CallBackResponse{}, nil
}
