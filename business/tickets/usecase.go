package tickets

import (
	"context"
	"ticketing/helper/guid"
	"ticketing/helper/seats"
	"time"
)

type TicketsUsecase struct {
	ticketsRepository Repository
	contextTimeout time.Duration
}

func NewTicketsUsecase(tr Repository, timeout time.Duration) Usecase {
	return &TicketsUsecase{
		ticketsRepository: tr,
		contextTimeout: timeout,
	}
}

func (tu *TicketsUsecase) Store(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	userDomain.BookingCode = guid.GenerateUUID()
	userDomain.Seats = seats.RandomString(2)
	err := tu.ticketsRepository.Store(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TicketsUsecase) Delete(ctx context.Context, id int) error {
	err := tu.ticketsRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TicketsUsecase) GetByID(ctx context.Context, id int) ([]Domain, error) {
	result, err := tu.ticketsRepository.GetByID(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

