package controllers

import (
	"ijash-marketplace/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []models.User

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil get data",
		Data:    users,
	})
}
