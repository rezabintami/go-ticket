package theater

import (
	"net/http"
	"strconv"
	"ticketing/business/theater"
	"ticketing/controllers/theater/request"
	"ticketing/controllers/theater/response"
	base_response "ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
)

type TheaterController struct {
	theaterUsecase theater.Usecase
}

func NewTheaterController(tu theater.Usecase) *TheaterController {
	return &TheaterController{
		theaterUsecase: tu,
	}
}

func (ctrl *TheaterController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Theater{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.theaterUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (ctrl *TheaterController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := ctrl.theaterUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (ctrl *TheaterController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()

	req := request.Theater{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.theaterUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Update Successfully")
}

func (ctrl *TheaterController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := ctrl.theaterUsecase.GetAll(ctx)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(result))
}
