package cUsers

import (
	"net/http"
	"strconv"
	"ticketing/configs"
	"ticketing/helper"
	"ticketing/models/users"
	"time"

	"github.com/labstack/echo/v4"
)

func UpdateProfile(ctx echo.Context) error {
	var usersInput users.UserInput
	ctx.Bind(&usersInput)

	var users users.Users
	users.Name = usersInput.Name
	users.Email = usersInput.Email
	users.Password = helper.HashGenerator(usersInput.Password)
	users.Language = usersInput.Language
	users.Updated_At = time.Now()
	result := configs.DB.Where("id = ?", usersInput.ID).Updates(&users).Error
	if result != nil {
		response := helper.APIResponse("failed update data", http.StatusBadRequest, "failed", nil)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	_ = configs.DB.Where("id = ?", usersInput.ID).First(&users).Error
	response := helper.APIResponse("success update data", http.StatusOK, "success", users)
	return ctx.JSON(http.StatusOK, response)
}

func GetProfile(ctx echo.Context) error {
	var users users.Users

	id, _ := strconv.Atoi(ctx.Param("id"))
	result := configs.DB.Where("id = ?", id).First(&users).Error
	if result != nil {
		response := helper.APIResponse("failed get data", http.StatusBadRequest, "failed", nil)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("success get data", http.StatusOK, "success", users)
	return ctx.JSON(http.StatusOK, response)
}
