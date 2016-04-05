package main

import "github.com/jinzhu/gorm"

// Quote stores the price of a Stock at a point in time
type Quote struct {
	gorm.Model
	StockID uint
	Stock   Stock
	Amount  float64
}

// SaveQuote saves a quote on a stock
func SaveQuote(db *gorm.DB, stock Stock, amount float64) Quote {
	quote := Quote{Stock: stock, Amount: amount}
	db.Create(&quote)
	return quote
}
