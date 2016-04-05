package main

import "testing"

func TestHoldingValue(t *testing.T) {
	s := Stock{Symbol: "FAKE", RowID: 10, ColID: 20}
	db.Create(&s)

	amount := 50.24
	SaveQuote(db, s, amount)

	// Value Value on Holding
	var holdingCount uint = 20
	holding := Holding{Stock: s, Count: holdingCount}

	expectedValue := float64(holdingCount) * amount
	value := holding.Value(db)
	if value != expectedValue {
		t.Error("Expected value", expectedValue, "got", value)
	}

	// Holding with No Stock
	invalidHolding := Holding{}
	value = invalidHolding.Value(db)
	if value != 0 {
		t.Error("Expected value 0, got", value)
	}
}
