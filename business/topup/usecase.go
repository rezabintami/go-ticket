package topup

import (
	"context"
	"ticketing/business/users"
	"time"
)

type topupUsecase struct {
	topupRepository Repository
	userRepository  users.Repository

	contextTimeout time.Duration
}

func NewTopUpUsecase(tr Repository, timeout time.Duration, us users.Repository) Usecase {
	return &topupUsecase{
		topupRepository: tr,
		contextTimeout:  timeout,
		userRepository:  us,
	}
}

func (tu *topupUsecase) Payment(ctx context.Context, topupDomain *Domain) error {
	err := tu.topupRepository.Payment(ctx, topupDomain)
	if err != nil {
		return err
	}
	result, err := tu.userRepository.GetByID(ctx, topupDomain.User_ID)
	if err != nil {
		return err
	}
	err = tu.userRepository.UpdateUser(ctx, &users.Domain{Balance: result.Balance + topupDomain.Balance}, topupDomain.User_ID)
	if err != nil {
		return err
	}

	return nil
}
