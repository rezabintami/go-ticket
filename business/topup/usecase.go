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

func (tu *TopupUsecase) CreateTransactions(ctx context.Context, paymentsDomain *payments.Domain, topupDomain *Domain, id int) (payments.DomainResponse, error) {
	//!MIDTRANS
	paymentsDomain.UserID = id
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

func (tu *TopupUsecase) Store(ctx context.Context, topupDomain *Domain) error {
	// err := tu.topupRepository.Store(ctx, topupDomain)
	// if err != nil {
	// 	return err
	// }
	// result, err := tu.userRepository.GetByID(ctx, topupDomain.UserID)
	// if err != nil {
	// 	return err
	// }
	// err = tu.userRepository.UpdateUser(ctx, &users.Domain{Amount: result.Amount + topupDomain.Amount}, topupDomain.UserID)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (tu *TopupUsecase) GetByID(ctx context.Context, id int) ([]Domain, error) {
	result, err := tu.topupRepository.GetByID(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
