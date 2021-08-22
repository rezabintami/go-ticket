package theater

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int           
	Name      string       
	Place     string         
	CreatedAt time.Time      
	UpdatedAt time.Time    
	DeletedAt gorm.DeletedAt 
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
