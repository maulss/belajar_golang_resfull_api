package app

import (
	"belajar_golang_resrful_api/helper"
	"database/sql"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db
}
