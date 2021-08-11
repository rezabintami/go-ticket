package users

import (
	"context"
	"strings"
	"ticketing/app/middleware"
	"ticketing/business"
	"ticketing/helper/encrypt"
	"time"
)

type UserUseCase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) UseCase {
	return &UserUseCase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (Domain, error) {
	result, err := uc.userRepository.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	result.Token = uc.jwtAuth.GenerateToken(result.ID)
	return result, nil
}

func (uc *UserUseCase) Register(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByEmail(ctx, userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return business.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return business.ErrInternalServer
	}
	err = uc.userRepository.Register(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}
