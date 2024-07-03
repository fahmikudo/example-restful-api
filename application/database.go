package application

import (
	"database/sql"
	"github.com/fahmikudo/example-restful-api/helper"
	"time"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:falcon21@tcp(127.0.0.1:3306)/belajar_golang_restful_api")
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
