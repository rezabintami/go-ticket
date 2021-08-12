package topup

import (
	"fmt"
	"net/http"
	"ticketing/business/topup"
	"ticketing/controllers/topup/request"
	"ticketing/helper/response"

	echo "github.com/labstack/echo/v4"
)

type TopUpController struct {
	topupUseCase topup.Usecase
}

func NewTopUpController(tc topup.Usecase) *TopUpController {
	return &TopUpController{
		topupUseCase: tc,
	}
}

func (ctrl *TopUpController) PaymentTopUp(c echo.Context) error {
	ctx := c.Request().Context()
	fmt.Println("start payment")

	req := request.TopUp{}
	if err := c.Bind(&req); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	fmt.Println("middle payment")
	err := ctrl.topupUseCase.Payment(ctx, req.ToDomain())
	if err != nil {
		return response.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	fmt.Println("finish payment")
	return response.NewSuccessResponse(c, "Successfully topup")
}
