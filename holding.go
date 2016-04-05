package main

import "github.com/jinzhu/gorm"

// Holding stores the number of stocks owned
type Holding struct {
	Stock Stock
	Count uint
}

// Value of Stocks in holding
func (h *Holding) Value(db *gorm.DB) float64 {
	return float64(h.Count) * h.Stock.Value(db)
}
