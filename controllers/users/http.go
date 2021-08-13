package users

import (
	"net/http"
	"ticketing/app/middleware"
	"ticketing/business/users"
	"ticketing/controllers/users/request"
	"ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.userUseCase.Register(ctx, req.ToDomain())
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, "Successfully inserted")
}

func (controller *UserController) Login(c echo.Context) error {
	var userLogin request.Users
	if err := c.Bind(&userLogin); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := controller.userUseCase.Login(c.Request().Context(), userLogin.Email, userLogin.Password)

	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, user)
}

func (controller *UserController) GetProfile(c echo.Context) error {
	id := middleware.GetUserId(c)
	// id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.userUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, user)
}

func (controller *UserController) UpdateProfile(c echo.Context) error {
	id := middleware.GetUserId(c)
	// id, _ := strconv.Atoi(c.Param("id"))
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := controller.userUseCase.UpdateUser(c.Request().Context(), req.ToDomain(), id)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	user, err := controller.userUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.NewSuccessResponse(c, user)
}
