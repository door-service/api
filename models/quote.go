package models

import "time"

// Quote is an estimate of all the parts, equipment, and labor for any given job
type Quote struct {
	Parts            Parts            `json:"parts"`
	EquipmentRentals EquipmentRentals `json:"equipment-rentals"`
	Labor            Labor            `json:"labor"`
}

// Quotes is a collection of individual quotes and the date they were created and updated
type Quotes []struct {
	Quote       Quote     `json:"quote"`
	DateCreated time.Time `json:"date-created"`
	DateUpdated time.Time `json:"date-updated"`
}
