package topup

import (
	"net/http"
	"ticketing/app/middleware"
	"ticketing/business/payments"
	"ticketing/business/topup"
	"ticketing/controllers/topup/request"
	"ticketing/controllers/topup/response"
	base_response "ticketing/helper/response"

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

func (ctrl *TopUpController) CreateTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	// req := request.TopUp{}
	// if err := c.Bind(&req); err != nil {
	// 	return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	// }

	resp,err := ctrl.topupUsecase.CreateTransactions(ctx, &payments.Domain{})
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, resp)
}

func (ctrl *TopUpController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.TopUp{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.topupUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully insert topup")
}

func (ctrl *TopUpController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := middleware.GetUserId(c)
	result, err := ctrl.topupUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(result))
}