package main

import (
	"github.com/joho/godotenv"
	"go_shop/internal/app"
	"log"
)

func main() {
	loadEnv()
	app := app.NewApp()
	app.Run()
	defer app.Shutdown()

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
