package models

// Labor represents the cost and number of men per quote
type Labor struct {
	Hours            float64 `json:"hours"`
	Workers          int     `json:"workers"`
	Rate             float64 `json:"rate"`
	IsPrevailingWage bool    `json:"is-prevailing-wage"`
}
