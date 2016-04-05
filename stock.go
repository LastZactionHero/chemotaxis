package main

import "github.com/jinzhu/gorm"

// Stock represents a single stock symbol
type Stock struct {
	gorm.Model
	Symbol string
	RowID  uint
	ColID  uint
}

// Value latest quote amount
func (s *Stock) Value(db *gorm.DB) float64 {
	var quote Quote
	db.Order("id DESC").Where("stock_id = ?", s.ID).First(&quote)
	return quote.Amount
}

// FindStock returns a stock by symbol, nil if not found
func FindStock(db *gorm.DB, symbol string) *Stock {
	var stock Stock
	if db.Where("symbol = ?", symbol).First(&stock).RecordNotFound() {
		return nil
	}
	return &stock
}
