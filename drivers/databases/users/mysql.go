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

func (repository *mysqlUsersRepository) GetByID(ctx context.Context, id int) (users.UserDomain, error) {
	usersById := Users{}
	result := repository.Conn.Where("id = ?", id).First(&usersById)
	if result.Error != nil {
		return users.UserDomain{}, result.Error
	}
	return usersById.toUserDomain(), nil
}

func (repository *mysqlUsersRepository) UpdateUser(ctx context.Context, userDomain *users.Domain, id int) error {
	usersUpdate := Users{}
	usersUpdate.Name = userDomain.Name
	usersUpdate.Email = userDomain.Email
	usersUpdate.Language = userDomain.Language
	usersUpdate.Balance = userDomain.Balance
	result := repository.Conn.Where("id = ?", id).Updates(&usersUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlUsersRepository) Login(ctx context.Context, id int) (users.Domain, error) {
	userLogin := Users{}
	result := repository.Conn.Where("id = ?", id).First(&userLogin)
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
