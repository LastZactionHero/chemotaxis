package main

import "github.com/jinzhu/gorm"

// Stock represents a single stock symbol
type Stock struct {
	gorm.Model
	Symbol string
	RowID  uint
	ColID  uint
}

// VALUE - returns latest quote
