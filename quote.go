package main

import "github.com/jinzhu/gorm"

// Quote stores the price of a Stock at a point in time
type Quote struct {
	gorm.Model
	StockID int
	Stock   Stock
	amount  float64
}
