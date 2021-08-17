package theater

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Place     string         `json:"place"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, data *Domain, id int) error
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, data *Domain, id int) error
	GetAll(ctx context.Context) ([]Domain, error)
}
