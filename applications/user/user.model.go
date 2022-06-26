package users

type UserModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive"`
}
