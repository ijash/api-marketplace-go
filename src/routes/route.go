package routes

import (
	"ijash-jwt-auth/src/controllers"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {

	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	// Define routes that require JWT authentication
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("PRIVATE_KEY_JWT"))))
	eAuth.GET("/me", controllers.GetUserProfileController)

	// Add a catch-all route for unmatched routes without authentication
	e.Any("/*", controllers.NotFoundController)
}
