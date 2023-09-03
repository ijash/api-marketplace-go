package controllers

import (
	"ijash-jwt-auth/src/helpers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotFoundController(c echo.Context) error {
	// Your custom "Not Found" response logic here
	log.Println("NotFoundController accessed")
	return c.JSON(http.StatusNotFound, helpers.NotFound())
}
