package topup

import (
	"net/http"
	"ticketing/app/middleware"
	"ticketing/business/topup"
	"ticketing/controllers/topup/request"
	"ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
)

type TopUpController struct {
	topupUsecase topup.Usecase
}

func NewTopUpController(tc topup.Usecase) *TopUpController {
	return &TopUpController{
		topupUsecase: tc,
	}
}

func (ctrl *TopUpController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.TopUp{}
	if err := c.Bind(&req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.topupUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, "Successfully topup")
}

func (ctrl *TopUpController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUserId(c)
	result, err := ctrl.topupUsecase.GetByID(ctx, id)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return response.NewSuccessResponse(c, result)
}