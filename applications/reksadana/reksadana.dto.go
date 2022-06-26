package reksadana

type CreateRequest struct {
	Name   string `json:"name" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
	Return int    `json:"return" validate:"required"`
}

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Model  `json:"data"`
}

type ResponseMany struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Data   *[]Model `json:"data"`
}

type ResponseError struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
