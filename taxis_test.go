package main

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	db = loadDatabase(DBFilenameTest)
	seedData(db)

	m.Run()

	os.Remove(DBFilenameTest)
}
