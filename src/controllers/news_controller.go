package controllers

import (
	"ijash-marketplace/src/configs"
	"ijash-marketplace/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddNewsController(c echo.Context) error {
	var newsRequest models.News

	c.Bind(&newsRequest)

	result := configs.DB.Create(&newsRequest)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "Success insert data to database",
		Data:    newsRequest,
	})
}

func GetDetailNewsController(c echo.Context) error {
	var id = c.Param("id")
	var data = models.News{}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: id,
		Data:    data,
	})
}

func GetNewsController(c echo.Context) error {
	var data []models.News

	result := configs.DB.Find(&data)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed get data from news",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Success get Data",
		Data:    data,
	})
}
