package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

const sampleQuoteCSVFilePath = "./data/sample_quotes_sp500.csv"

func loadQuoteSample(db *gorm.DB) {
	f, _ := os.Open(sampleQuoteCSVFilePath)
	r := csv.NewReader(bufio.NewReader(f))
	quotes, _ := r.ReadAll()

	for _, row := range quotes {
		symbol := row[0]
		amount, _ := strconv.ParseFloat(row[1], 64)

		stock := FindStock(db, symbol)
		SaveQuote(db, *stock, amount)
	}
}
