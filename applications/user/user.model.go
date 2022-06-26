package users

type Model struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive"`
}
