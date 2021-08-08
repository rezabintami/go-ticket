package auth

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type RegisterInput struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password,omitempty"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
	Language string  `json:"language"`
}
