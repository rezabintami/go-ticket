package topup

import (
	"context"
	"ticketing/business/payments"
	"ticketing/business/users"
	"time"
)

type TopupUsecase struct {
	topupRepository Repository
	userRepository  users.Repository
	paymentsRepository payments.Repository
	contextTimeout time.Duration
}

func NewTopUpUsecase(tr Repository, timeout time.Duration, us users.Repository, pay payments.Repository) Usecase {
	return &TopupUsecase{
		topupRepository: tr,
		contextTimeout:  timeout,
		userRepository:  us,
		paymentsRepository: pay,
	}
}

func (tu *TopupUsecase) CreateTransactions(ctx context.Context, paymentsDomain *payments.Domain) error {
	//!MIDTRANS
	tu.paymentsRepository.Transactions(ctx, paymentsDomain)

	return nil
}



func (tu *TopupUsecase) Store(ctx context.Context, topupDomain *Domain) error {
	//!MIDTRANS

	err := tu.topupRepository.Store(ctx, topupDomain)
	if err != nil {
		return err
	}
	result, err := tu.userRepository.GetByID(ctx, topupDomain.UserID)
	if err != nil {
		return err
	}
	err = tu.userRepository.UpdateUser(ctx, &users.Domain{Balance: result.Balance + topupDomain.Balance}, topupDomain.UserID)
	if err != nil {
		return err
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

