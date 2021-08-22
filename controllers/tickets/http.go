package tickets

import (
	"net/http"
	"strconv"
	"ticketing/app/middleware"
	"ticketing/business/tickets"
	"ticketing/controllers/tickets/request"
	"ticketing/controllers/tickets/response"
	base_response "ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
)

type TicketsController struct {
	ticketsUsecase tickets.Usecase
}

func NewTicketsController(tu tickets.Usecase) *TicketsController {
	return &TicketsController{
		ticketsUsecase: tu,
	}
}

func (ctrl *TicketsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUserId(c)
	req := request.Tickets{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.ticketsUsecase.Store(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully create ticket")
}

func (ctrl *TicketsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))
	userId := middleware.GetUserId(c)
	req := request.Tickets{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.ticketsUsecase.Delete(ctx, id,userId)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (ctrl *TicketsController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUserId(c)
	result, err := ctrl.ticketsUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(result))
}
