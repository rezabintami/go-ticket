package movies

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"ticketing/business/movies"
	"ticketing/controllers/movies/response"
	base_response "ticketing/helper/response"

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
		return base_response.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required search"))
	}
	result, err := ctrl.moviesUsecase.Fetch(ctx, url.QueryEscape(search), search)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromAllDomain(result))
}

func (ctrl *MovieController) GetDetailMovies(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movie, err := ctrl.moviesUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(movie))
}
