package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Stock represents a single stock symbol
type Stock struct {
	gorm.Model
	Symbol string
	RowID  uint
	ColID  uint
}

var db *gorm.DB

func main() {
	db = loadDatabase()

	loadGradient()
	printGradient()
}

func loadDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	if db.HasTable("stocks") {
		db.DropTable("stocks")
	}
	db.CreateTable(&Stock{})
	return db
}

func loadGradient() [][]string {
	f, _ := os.Open("./data/sp500_grid.csv")
	r := csv.NewReader(bufio.NewReader(f))
	gradient, _ := r.ReadAll()
	for rowID, row := range gradient {
		for colID, symbol := range row {
			db.Create(&Stock{Symbol: symbol, RowID: uint(rowID), ColID: uint(colID)})
		}
	}
	return gradient
}

func printGradient() {
	var stocks []Stock
	db.Order("row_id, col_id").Find(&stocks)

	var lastRowID uint
	for _, stock := range stocks {
		if stock.RowID > lastRowID {
			lastRowID = stock.RowID
			fmt.Printf("\n")
		}

		fmt.Printf(stock.Symbol)
		if stock.RowID > lastRowID {
			lastRowID = stock.RowID
			fmt.Printf("\n")
		} else {
			fmt.Printf("\t| ")
		}
	}
	fmt.Printf("\n\n")
}

// stock
//   symbol
//   sector
//   value
//   row_id
//   col_id
//
// quote
//   amount
//   datetime
//   stock_id
