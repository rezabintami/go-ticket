package response

import (
	"ticketing/business/users"
	"time"
)

type Users struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password,omitempty"`
	Email      string    `json:"email"`
	Balance    float64   `json:"balance"`
	Language   string    `json:"language"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func FromDomain(userDomain users.Domain) Users {
	return Users{
		ID:         userDomain.ID,
		Name:       userDomain.Name,
		Password:   userDomain.Password,
		Email:      userDomain.Email,
		Language:   userDomain.Language,
		Created_At: userDomain.Created_At,
		Updated_At: userDomain.Updated_At,
	}
}
