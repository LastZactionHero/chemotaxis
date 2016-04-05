package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GradientCSVFilePath path to gradient stock grid csv
const GradientCSVFilePath = "./data/sp500_grid.csv"

// DBFilename SQLite Database filename
const DBFilename = "db/development.db"

// DBFilenameTest SQLite Database filename
const DBFilenameTest = "db/test.db"

// var db *gorm.DB
var gradient Gradient

func main() {
	db := loadDatabase(DBFilename)

	seedData(db)

	gradient.Load(db)
	gradient.Print()

	loadQuoteSample(db)

	var stock Stock
	db.Where("symbol = ?", "AAPL").First(&stock)
	fmt.Println(stock.Symbol)
	fmt.Println(stock.Value(db))
}

func loadDatabase(filename string) *gorm.DB {
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic("failed to connect database")
	}

	clearTables := []string{"stocks", "quotes"}
	for _, table := range clearTables {
		if db.HasTable(table) {
			db.DropTable(table)
		}
	}

	db.CreateTable(&Stock{})
	db.CreateTable(&Quote{})
	return db
}

func seedData(db *gorm.DB) [][]string {
	f, _ := os.Open(GradientCSVFilePath)
	r := csv.NewReader(bufio.NewReader(f))
	gradient, _ := r.ReadAll()
	for rowID, row := range gradient {
		for colID, symbol := range row {
			db.Create(&Stock{Symbol: symbol, RowID: uint(rowID), ColID: uint(colID)})
		}
	}
	return gradient
}
