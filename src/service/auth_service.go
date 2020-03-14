package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"

	"github.com/IkezawaYuki/pictweet-oauth/src/authpb"
	"github.com/IkezawaYuki/pictweet-oauth/src/infrastructure"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

// AuthService is ...
type AuthService struct {
	redisHandler infrastructure.RedisHandler
}

func NewAuthService(h infrastructure.RedisHandler) *AuthService {
	return &AuthService{
		redisHandler: h,
	}
}

// Login ログイン処理の実行　Googleのプロバイダを使用。
func (s *AuthService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	fmt.Println("Login function is invoked")
	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatalln("failed to get provider: ", provider, "-", err)
	}
	loginURL, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		log.Fatalln("error is occured when calling GetBeginAuthURL methods: ", provider, "-", err)
	}
	fmt.Println(loginURL)
	return &authpb.LoginResponse{
		RedirectUrl: loginURL,
	}, nil
}

// CallBack Googleでの認証後の処理
func (s *AuthService) CallBack(ctx context.Context, req *authpb.CallBackRequest) (*authpb.CallBackResponse, error) {
	fmt.Println("call back function is invoked")
	provider, err := gomniauth.Provider("google")
	if err != nil {
		log.Fatalln("failed to get provider", provider, "-", err)
	}
	creds, err := provider.CompleteAuth(objx.MustFromURLQuery("code=" + req.GetCode()))
	if err != nil {
		log.Fatalln("authentication could not be completed.")
	}
	user, err := provider.GetUser(creds)
	if err != nil {
		log.Fatalln("failed to get user", provider, "-", err)
	}
	fmt.Println(user)

	uuidObj, _ := uuid.NewUUID()
	token := uuidObj.String()
	email := user.Email()
	s.redisHandler.SetWithExpire(token, email)

	return &authpb.CallBackResponse{
		Token: uuidObj.String(),
	}, nil
}

// VerifyToken ログインしているかどうか
func (s *AuthService) VerifyToken(ctx context.Context, req *authpb.VerifyTokenRequest) (*authpb.VerifyTokenResponse, error) {
	email := s.redisHandler.Get(req.GetToken())
	if email == "" {
		return &authpb.VerifyTokenResponse{
			IsLoggedId: false,
		}, nil
	}
	return &authpb.VerifyTokenResponse{
		IsLoggedId: true,
	}, nil
}
