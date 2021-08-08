package users

type UserInput struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password,omitempty"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
	Language string  `json:"language"`
}
