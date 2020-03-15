package main

import (
	"github.com/IkezawaYuki/pictweet-oauth/src/infrastructure/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	server.Run()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env  %v", err)
	}
}
