package conventional_invoice

type Model struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Tenor  int    `json:"tenor"`
	Grade  string `json:"grade"`
	Rate   int    `json:"rate"`
}
