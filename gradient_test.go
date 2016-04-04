package main

import "testing"

func TestGradientLoad(t *testing.T) {
	var g Gradient
	g.Load(db)

	if len(g.Stocks) != 22 {
		t.Error("Expected 22 rows, got ", len(g.Stocks))
	}
	for rowIdx, row := range g.Stocks {
		if len(row) != 23 {
			t.Error("Expected 23 stocks in each row, got", len(row), "in row", rowIdx)
		}
	}
}
