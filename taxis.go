package main

import (
	"bufio"
	"encoding/csv"
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
}

func loadDatabase(filename string) *gorm.DB {
	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		panic("failed to connect database")
	}

	if db.HasTable("stocks") {
		db.DropTable("stocks")
	}
	db.CreateTable(&Stock{})
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
