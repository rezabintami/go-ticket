package movies

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"ticketing/business/movies"
	"ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
)

type MovieController struct {
	moviesUsecase movies.Usecase
}

func NewMovieController(mu movies.Usecase) *MovieController {
	return &MovieController{
		moviesUsecase: mu,
	}
}

func (ctrl *MovieController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()
	search := c.QueryParam("search")
	if strings.TrimSpace(search) == "" {
		return response.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required search"))
	}
	err := ctrl.moviesUsecase.Store(ctx, url.QueryEscape(search))
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, "Successfully inserted")
}
