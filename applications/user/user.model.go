package users

import (
	"dompet/config"

	"github.com/golang-jwt/jwt"
)

type Model struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive"`
}

func (u *Model) GenerateToken(others map[string]interface{}) (string, error) {
	jwtKey := config.Config.JWT.Secret

	data := jwt.MapClaims{
		"id": u.Email,
	}

	for k, v := range others {
		data[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
