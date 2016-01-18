package models

// Part represents a part needed for the quote
type Part struct {
	Manufacturer Manufacturer `json:"manufacturer"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	UnitPrice    float64      `json:"unit-price"`
}

// Parts is a collection of individual parts and their associated quantity
type Parts []struct {
	Part     Part `json:"part"`
	Quantity int  `json:"quantity"`
}
