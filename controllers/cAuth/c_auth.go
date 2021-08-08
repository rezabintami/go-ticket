package cAuth

import (
	"net/http"
	"ticketing/helper"
	"ticketing/models/auth"

	"github.com/labstack/echo"
)

func Login(ctx echo.Context) error {
	var auth auth.Auth
	ctx.Bind(&auth)
	response := helper.APIResponse("Login Success", http.StatusOK, "success", auth)

	return ctx.JSON(http.StatusOK, response)
}

func Register() {
	var auth auth.Auth
	ctx.Bind(&auth)
	response := helper.APIResponse("Login Success", http.StatusOK, "success", auth)

	return ctx.JSON(http.StatusOK, response)
}
