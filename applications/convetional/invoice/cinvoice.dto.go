package conventional_invoice

type CreateRequest struct {
	Name   string `json:"name" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
	Tenor  int    `json:"tenor" validate:"required"`
	Grade  string `json:"grade" validate:"required"`
	Rate   int    `json:"rate" validate:"required"`
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
