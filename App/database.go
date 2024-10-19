package app

import (
	"database/sql"
	"fmt"
	"restfull-api/helper"
	"time"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

func NewDb() *sql.DB {

	const (
		user = "postgres"

		dbName   = "CRUD"
		password = "admin123"
	)
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)

	db.SetMaxOpenConns(25)           // Maksimal 25 koneksi yang terbuka
	db.SetMaxIdleConns(25)           // Maksimal 25 koneksi idle
	db.SetConnMaxLifetime(time.Hour) // Maksimal waktu hidup 1 jam per koneksi

	if err != nil {
		fmt.Println("Errrooorrr ")
		helper.PanicIfErr(err)
	}

	return db
}
