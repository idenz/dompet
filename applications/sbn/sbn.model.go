package sbn

type Model struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Tenor  int    `json:"tenor"`
	Rate   int    `json:"rate"`
	Type   string `json:"type"`
}
