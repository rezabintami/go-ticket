package users

import (
	"net/http"
	"ticketing/app/middleware"
	"ticketing/business/users"
	"ticketing/controllers/users/request"
	"ticketing/controllers/users/response"
	base_response "ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.userUsecase.Register(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var userLogin request.Users
	if err := c.Bind(&userLogin); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := controller.userUsecase.Login(ctx, userLogin.Email, userLogin.Password)

	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result := struct {
		Token string `json:"token"`
	}{Token: token}
	return base_response.NewSuccessResponse(c, result)
}

func (controller *UserController) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUserId(c)
	user, err := controller.userUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(user))
}

func (controller *UserController) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUserId(c)
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := controller.userUsecase.UpdateUser(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	user, err := controller.userUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessResponse(c, response.FromDomain(user))
}
