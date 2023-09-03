package main

import (
	"ijash-jwt-auth/src/configs"
	"ijash-jwt-auth/src/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(os.Getenv("PORT"))

}
