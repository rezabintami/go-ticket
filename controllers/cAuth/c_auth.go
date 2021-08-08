package cAuth

import (
	"net/http"
	"ticketing/configs"
	"ticketing/helper"
	"ticketing/models/auth"
	"ticketing/models/users"

	"github.com/labstack/echo/v4"
)

func Login(ctx echo.Context) error {
	var auth auth.LoginInput
	ctx.Bind(&auth)
	response := helper.APIResponse("login success", http.StatusOK, "success", auth)

	return ctx.JSON(http.StatusOK, response)
}

func Register(ctx echo.Context) error {
	var register auth.RegisterInput
	ctx.Bind(&register)

	var users users.Users
	users.Name = register.Name
	users.Email = register.Email
	users.Password = helper.HashGenerator(register.Password)
	users.Balance = register.Balance
	users.Language = register.Language

	result := configs.DB.Create(&users)
	if result.Error != nil {
		response := helper.APIResponse("register failed", http.StatusInternalServerError, "failed", nil)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := helper.APIResponse("register success", http.StatusOK, "success", users)
	return ctx.JSON(http.StatusOK, response)
}
