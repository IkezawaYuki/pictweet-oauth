package server

import (
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

// StartApp サーバーの起動
func StartApp() {
	gomniauth.SetSecurityKey(os.Getenv("セキュリティーキー"))
	gomniauth.WithProviders(
		google.New(os.Getenv("クライアントID"), "秘密の値", "http://localhost:8080/auth/callback/google"),
	)

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

	}
	defer c.Close()

}
