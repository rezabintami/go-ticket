package movies

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
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
	result, err := ctrl.moviesUsecase.Fetch(ctx, url.QueryEscape(search), search)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, result)
}

func (ctrl *MovieController) GetDetailMovies(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movie, err := ctrl.moviesUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return response.NewSuccessResponse(c, movie)
}
