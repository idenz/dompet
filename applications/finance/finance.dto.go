package finances

type CreateRequest struct {
	Name  string `json:"name" validate:"required"`
	Count int    `json:"count"`
	Sub   string `json:"sub"`
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
