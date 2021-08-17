package tickets

import (
	"context"
	"ticketing/business/users"
	"ticketing/helper/guid"
	"ticketing/helper/seats"
	"time"
)

type TicketsUsecase struct {
	ticketsRepository Repository
	userRepository    users.Repository
	contextTimeout    time.Duration
}

func NewTicketsUsecase(tr Repository, us users.Repository, timeout time.Duration) Usecase {
	return &TicketsUsecase{
		ticketsRepository: tr,
		userRepository:    us,
		contextTimeout:    timeout,
	}
}

func (tu *TicketsUsecase) Store(ctx context.Context, ticketDomain *Domain, id int) error {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	ticketDomain.BookingCode = guid.GenerateUUID()
	ticketDomain.Seats = seats.RandomString(2)
	err := tu.ticketsRepository.Store(ctx, ticketDomain)
	if err != nil {
		return err
	}

	result, err := tu.userRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = tu.userRepository.UpdateUser(ctx, &users.Domain{Balance: result.Balance - ticketDomain.TotalPrice}, id)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TicketsUsecase) Delete(ctx context.Context, id int, userId int) error {
	res, err := tu.ticketsRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = tu.ticketsRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	result, err := tu.userRepository.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	err = tu.userRepository.UpdateUser(ctx, &users.Domain{Balance: result.Balance + (res.TotalPrice * 0.9)}, userId)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TicketsUsecase) GetByID(ctx context.Context, id int) ([]Domain, error) {
	
	result, err := tu.ticketsRepository.GetAllByID(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
