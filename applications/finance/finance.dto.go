package finances

type Data struct {
	Model
}

type CreateRequest struct {
	Name  string `json:"name" validate:"required"`
	Count int    `json:"count"`
	Sub   string `json:"sub"`
}

type Response struct {
	Data
}