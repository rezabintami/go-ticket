package tickets

import (
	"context"
	"ticketing/business/tickets"

	"gorm.io/gorm"
)

type mysqlTicketsRepository struct {
	Conn *gorm.DB
}

func NewMySQLTicketsRepository(conn *gorm.DB) tickets.Repository {
	return &mysqlTicketsRepository{
		Conn: conn,
	}
}

func (repository *mysqlTicketsRepository) Store(ctx context.Context, ticketsDomain *tickets.Domain) error {
	rec := fromDomain(*ticketsDomain)
	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTicketsRepository) Delete(ctx context.Context, id int) error {
	tickets := Tickets{}
	result := repository.Conn.Where("id = ?", id).Delete(&tickets)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTicketsRepository) GetByID(ctx context.Context, id int) (tickets.Domain, error) {
	tic := Tickets{}
	result := repository.Conn.Where("id = ?", id).Find(&tic)
	if result.Error != nil {
		return tickets.Domain{}, result.Error
	}
	return tic.toDomain(), nil
}

func (repository *mysqlTicketsRepository) GetAllByID(ctx context.Context, id int) ([]tickets.Domain, error) {
	tic := []Tickets{}
	result := repository.Conn.Where("user_id = ?", id).Find(&tic)
	if result.Error != nil {
		return []tickets.Domain{}, result.Error
	}
	var allTickets []tickets.Domain
	for _, value := range tic {
		allTickets = append(allTickets, value.toDomain())
	}
	return allTickets, nil
}