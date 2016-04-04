package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func printCSVUrl(db *gorm.DB) {
	var stocks []Stock
	db.Order("symbol").Find(&stocks)
	var url = "http://download.finance.yahoo.com/d/quotes.csv?s="

	symbols := make([]string, len(stocks))
	for stockIdx, stock := range stocks {
		symbols[stockIdx] = stock.Symbol
	}
	url += strings.Join(symbols, "+")
	url += "&f=sl1d1t1c1ohgv&e=.csv"

	fmt.Println(url)
}
