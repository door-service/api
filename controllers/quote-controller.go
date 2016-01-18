package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/door-service/api/models"
)

// QuoteController is a controller for quotes
type QuoteController struct{}

// NewQuoteController constructs an empty Quotecontroller
func NewQuoteController() *QuoteController {
	return &QuoteController{}
}

// GetAll retrieves all of the quotes that exist
func (qc QuoteController) GetAll(w http.ResponseWriter, r *http.Request) {

	q := models.Quotes{
		{
			Quote: models.Quote{
				Parts: models.Parts{
					{
						Part: models.Part{
							Manufacturer: models.Manufacturer{
								Name:    "Manny Inc.",
								Address: "Tes Drive Ln.",
								Phone:   "555-5555",
								Email:   "test@email.com",
							},
							Name:        "Part No. 1",
							Description: "The first of the parts",
							UnitPrice:   0.0,
						},
						Quantity: 0,
					},
				},
				EquipmentRentals: models.EquipmentRentals{
					{
						Equipment: models.Equipment{
							Name:          "Equip No. 1",
							Description:   "The first of the equipment",
							PerDayPrice:   0.0,
							DeliveryPrice: 0.0,
						},
						Days: 0,
					},
				},
				Labor: models.Labor{
					Hours:            12.5,
					Workers:          2,
					Rate:             32.00,
					IsPrevailingWage: false,
				},
			},
			DateCreated: time.Now(),
			DateUpdated: time.Now(),
		},
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(q); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
