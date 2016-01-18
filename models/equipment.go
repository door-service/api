package models

// Equipment represents the equipment needed for the quote
type Equipment struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	PerDayPrice   float64 `json:"per-day-price"`
	DeliveryPrice float64 `json:"delivery-price"`
}

// EquipmentRentals is a collection of equipment and the amount of days they are to be rented for
type EquipmentRentals []struct {
	Equipment Equipment `json:"equipment"`
	Days      int       `json:"days"`
}
