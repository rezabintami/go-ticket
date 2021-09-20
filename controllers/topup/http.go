package topup

import (
	"net/http"
	"ticketing/app/middleware"
	"ticketing/business/topup"
	"ticketing/controllers/topup/request"
	"ticketing/controllers/topup/response"
	"ticketing/helper/guid"
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
	id := middleware.GetUserId(c)

	req := request.TopUp{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	req.OrderID = guid.GenerateUUID()

	resp, err := ctrl.topupUsecase.CreateTransactions(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromPaymentDomain(resp))
}

func (ctrl *TopUpController) TransactionCallbackHandler(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.MidtransCallback{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	
	err := ctrl.topupUsecase.Update(ctx, req.HandlerToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Successfully")
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
