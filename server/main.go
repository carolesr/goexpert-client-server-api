package main

import (
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := setupDB()
	if err != nil {
		panic(err)
	}

	service := Service{
		db: db,
	}

	handler := Handler{
		service: &service,
	}

	http.HandleFunc("/cotacao", handler.ExchangeRateHandler)
	http.ListenAndServe(":8080", nil)
}

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Model{})

	return db, nil
}
