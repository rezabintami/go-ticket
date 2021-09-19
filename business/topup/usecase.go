package topup

import (
	"context"
	"ticketing/business/payments"
	"ticketing/business/users"
	"time"
)

type TopupUsecase struct {
	topupRepository    Repository
	userRepository     users.Repository
	paymentsRepository payments.Repository
	contextTimeout     time.Duration
}

func NewTopUpUsecase(tr Repository, timeout time.Duration, us users.Repository, pay payments.Repository) Usecase {
	return &TopupUsecase{
		topupRepository:    tr,
		contextTimeout:     timeout,
		userRepository:     us,
		paymentsRepository: pay,
	}
}

func (tu *TopupUsecase) CreateTransactions(ctx context.Context, topupDomain *Domain, id int) (payments.DomainResponse, error) {
	//!MIDTRANS
	topupDomain.UserID = id

	result, err := tu.topupRepository.Store(ctx, topupDomain)
	if err != nil {
		return payments.DomainResponse{}, err
	}

	response, err := tu.paymentsRepository.Transactions(ctx, &result)
	if err != nil {
		return payments.DomainResponse{}, err
	}

	return response, nil
}

func (tu *TopupUsecase) Update(ctx context.Context, topupDomain *Domain) error {
	if topupDomain.PaymentName == "credit_card" && topupDomain.Status == "capture" && topupDomain.FraudStatus == "accept" {
		topupDomain.Status = "paid"
	} else if topupDomain.Status == "settlement" {
		topupDomain.Status = "paid"
	} else if topupDomain.Status == "deny" || topupDomain.Status == "expire" || topupDomain.Status == "cancel" {
		topupDomain.Status = "canceled"
	}

	err := tu.topupRepository.Update(ctx, topupDomain)
	if err != nil {
		return err
	}
	
	if topupDomain.Status == "paid" {
		result, err := tu.userRepository.GetByID(ctx, topupDomain.UserID)
		if err != nil {
			return err
		}

		err = tu.userRepository.UpdateUser(ctx, &users.Domain{Amount: result.Amount + topupDomain.Amount}, topupDomain.UserID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (tu *TopupUsecase) GetByID(ctx context.Context, id int) ([]Domain, error) {
	result, err := tu.topupRepository.GetByID(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
