package main

import (
	"fmt"
	"github.com/IkezawaYuki/pictweet-oauth/src/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"os"
)

func main() {
	gomniauth.SetSecurityKey("セキュリティキー？")
	gomniauth.WithProviders(
		google.New(os.Getenv("CLIENT_ID"), os.Getenv("SECRET_VALUE"), "http://localhost:8080/auth/callback/google"),
	)

	go server.StartGrpcServer()
	server.StartHttpServer()
}

func init(){
	fmt.Println("init...")
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		logrus.Fatalf("Error loading .env")
	}
}

