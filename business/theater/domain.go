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
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, data *Domain, id int) error
	GetAll(ctx context.Context, data *Domain) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, data *Domain, id int) error
	GetAll(ctx context.Context, data *Domain) ([]Domain, error)
}
