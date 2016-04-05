package main

import "testing"

func TestSaveQuote(t *testing.T) {
	stock := Stock{Symbol: "FAKE", RowID: 10, ColID: 20}
	db.Create(&stock)

	amount := 99.99
	SaveQuote(db, stock, amount)

	var quote Quote
	db.Where("stock_id = ?", stock.ID).First(&quote)

	if quote.ID <= 0 {
		t.Error("Quote not saved")
	} else if quote.Amount != amount {
		t.Error("Expected amount", amount, "got", quote.Amount)
	}
}
