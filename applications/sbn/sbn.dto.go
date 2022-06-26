package sbn

type CreateRequest struct {
	Name   string `json:"name" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
	Tenor  int    `json:"tenor" validate:"required"`
	Rate   int    `json:"rate"`
	Type   string `json:"type" validate:"required"`
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
