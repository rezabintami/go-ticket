package users

import (
	"context"
	"ticketing/business/users"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySQLUserRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (repository *mysqlUsersRepository) Login(ctx context.Context, email, password string) (users.Domain, error) {
	userLogin := Users{}
	result := repository.Conn.Where("email = ? AND password = ?", email, password).First(&userLogin)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return userLogin.toDomain(), nil
}

func (nr *mysqlUsersRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlUsersRepository) Register(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
