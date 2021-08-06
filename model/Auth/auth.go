package auth

import "time"

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Register struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"password,omitempty"`
	Email      string    `json:"email"`
	Balance    float64   `json:"balance"`
	Language   string    `json:"language"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
