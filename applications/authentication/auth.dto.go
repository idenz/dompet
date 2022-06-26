package authentication

import users "dompet/applications/user"

type Data struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseRegister struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   users.Model `json:"data"`
}

type ResponseLogin struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type ResponseError struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
