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

func TestFindStock(t *testing.T) {
	s := Stock{Symbol: "FAKE", RowID: 10, ColID: 99}
	db.Create(&s)

	found := FindStock(db, s.Symbol)
	if found == nil {
		t.Error("Expected found to exist")
	} else if found.Symbol != s.Symbol {
		t.Error("Expected stock", s.Symbol, "got", found.Symbol)
	}

	missing := FindStock(db, "GRBG")
	if missing != nil {
		t.Error("Expecting stock to be nil, got", missing)
	}
}
