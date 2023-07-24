package app

import (
	"database/sql"
	"time"

	"github.com/dickidarmawansaputra/belajar-go-database-migration/helper"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "dickids:rahasia@tcp(localhost:3306)/belajar-go-database-migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
