package app

import (
	"database/sql"
	"fmt"
	"restfull-api/helper"

	_ "github.com/lib/pq"
)

func NewDb() *sql.DB {

	const (
		host     = "localhost"
		user     = "postgres"
		port     = 5432
		dbName   = "CRUD"
		password = "admin123"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Cek")
		helper.PanicIfErr(err)
	}

	return db
}
