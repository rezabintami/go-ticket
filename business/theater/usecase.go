package theater

import (
	"context"
	"time"
)

type TheaterUsecase struct {
	theaterRepository Repository
	contextTimeout    time.Duration
}

func NewTheaterUsecase(tu Repository, timeout time.Duration) Usecase {
	return &TheaterUsecase{
		theaterRepository: tu,
		contextTimeout:    timeout,
	}
}

func (tu *TheaterUsecase) Store(ctx context.Context, theaterDomain *Domain) error {
	err := tu.theaterRepository.Store(ctx, theaterDomain)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TheaterUsecase) Delete(ctx context.Context, id int) error {
	err := tu.theaterRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TheaterUsecase) Update(ctx context.Context, theaterDomain *Domain, id int) error {
	err := tu.theaterRepository.Update(ctx, theaterDomain, id)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TheaterUsecase) GetAll(ctx context.Context, theaterDomain *Domain) ([]Domain, error) {
	result, err := tu.theaterRepository.GetAll(ctx, theaterDomain)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
