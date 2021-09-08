package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	_config "ticketing/app/config"
	"ticketing/app/middleware"
	"ticketing/business/users"
	"ticketing/controllers/users/request"
	"ticketing/controllers/users/response"
	base_response "ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     _config.GetConfig().Google.ClientID,
		ClientSecret: _config.GetConfig().Google.Secret,
		RedirectURL:  "http://localhost:8000/api/v1/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	randomstate = "random"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (controller *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.userUsecase.Register(ctx, req.ToDomain())
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

//! OAuth2 Google
func (controller *UserController) Google(c echo.Context) error {
	return c.Render(http.StatusOK, "googleauth.html", nil)
}

func (controller *UserController) LoginGoogle(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL(randomstate)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil
}

func (controller *UserController) HandleGoogle(c echo.Context) error {
	fmt.Println("masuk handle google")
	ctx := c.Request().Context()

	if randomstate != c.QueryParam("state") {
		return base_response.NewErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("invalid session state: %s", randomstate))
	}

	token, err := googleOauthConfig.Exchange(ctx, c.QueryParam("code"))
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	client := googleOauthConfig.Client(ctx, token)
	UserInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	defer UserInfo.Body.Close()

	req := request.Users{}
	json.NewDecoder(UserInfo.Body).Decode(&req)
	
	err = controller.userUsecase.Register(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}
