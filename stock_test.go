package main

import "testing"

func TestStockValue(t *testing.T) {
	s := Stock{Symbol: "FAKE", RowID: 10, ColID: 20}
	db.Create(&s)

	firstQuote := Quote{StockID: s.ID, Amount: 50.00}
	db.Create(&firstQuote)
	secondQuote := Quote{StockID: s.ID, Amount: 99.99}
	db.Create(&secondQuote)

	if s.Value(db) != secondQuote.Amount {
		t.Error("Expected value", secondQuote.Amount, "got", s.Value(db))
	}
}
