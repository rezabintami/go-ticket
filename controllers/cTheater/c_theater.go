package cTheater

import (
	"net/http"
	"ticketing/configs"
	"ticketing/helper"
	"ticketing/models/theater"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateTheater(ctx echo.Context) error {
	var theaterInput theater.TheaterInput
	ctx.Bind(&theaterInput)

	var theater theater.Theater
	theater.Name = theaterInput.Name
	theater.Place = theaterInput.Place
	theater.Created_At = time.Now()
	result := configs.DB.Create(&theater)
	if result.Error != nil {
		response := helper.APIResponse("create theater failed", http.StatusInternalServerError, "failed", nil)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := helper.APIResponse("create success", http.StatusOK, "success", theater)
	return ctx.JSON(http.StatusOK, response)
}

func GetTheater(ctx echo.Context) error {
	var theater []theater.Theater

	result := configs.DB.Find(&theater)
	if result.Error != nil {
		response := helper.APIResponse("failed get data", http.StatusInternalServerError, "failed", nil)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := helper.APIResponse("success get data", http.StatusOK, "success", theater)
	return ctx.JSON(http.StatusOK, response)
}
