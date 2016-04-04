package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Gradient represents a grid of Stocks
type Gradient struct {
	Stocks [][]Stock
}

// Load stocks from database into a matrix
func (g *Gradient) Load(db *gorm.DB) {
	var stocks []Stock
	db.Order("row_id, col_id").Find(&stocks)

	for _, stock := range stocks {
		if stock.RowID+1 > uint(len(g.Stocks)) {
			var stockRow []Stock
			g.Stocks = append(g.Stocks, stockRow)
		}
		g.Stocks[stock.RowID] = append(g.Stocks[stock.RowID], stock)
	}
}

// Print the gradient matrix
func (g *Gradient) Print() {
	for _, row := range g.Stocks {
		for _, stock := range row {
			fmt.Printf("%s\t| ", stock.Symbol)
		}
		fmt.Printf("\n")
	}
}
