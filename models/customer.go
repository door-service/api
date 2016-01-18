package models

// Customer represents the company that the quote belongs to
type Customer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Quotes  Quotes `json:"quotes"`
}
