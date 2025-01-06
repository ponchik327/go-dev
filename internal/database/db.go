package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) InitDB() {
	// в dsn вводим данные, которые мы указали при создании контейнера
	dsn := "host=localhost user=postgres password=12345678 dbname=postgres port=5432 sslmode=disable"
	var err error
	db.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}
