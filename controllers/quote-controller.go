package controllers

import "net/http"

// QuoteController is a controller for quotes
type QuoteController struct{}

// NewQuoteController constructs an empty Quotecontroller
func NewQuoteController() *QuoteController {
	return &QuoteController{}
}

// GetAll retrieves all of the quotes that exist
func (qc QuoteController) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)

	// if err := json.NewEncoder(w).Encode(q); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
