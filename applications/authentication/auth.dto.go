package authentication

import users "dompet/applications/user"

type Data struct {
	Token string `json:"token"`
	Email string `json:"email"`
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
