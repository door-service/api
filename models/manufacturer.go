package models

// Manufacturer represents the creator of the part
type Manufacturer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
