package server

import (
	"os"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.dip-net.co.jp/dip/robotics-auth/src/infrastructure"
)

// StartApp サーバーの起動
func StartApp() {
	gomniauth.SetSecurityKey(os.Getenv("セキュリティーキー"))
	gomniauth.WithProviders(
		google.New(os.Getenv("クライアントID"), "秘密の値", "http://localhost:8080/auth/callback/google"),
	)

	c := infrastructure.GetConnection()
	defer c.Close()

}
