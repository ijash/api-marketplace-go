package main

import (
	"fmt"
	"ijash-jwt-auth/src/configs"
	"ijash-jwt-auth/src/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// LoadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	// e.Start(":8000")
	e.Start(os.Getenv("PORT"))

}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("Environment: %s", os.Getenv("ENV"))
}
