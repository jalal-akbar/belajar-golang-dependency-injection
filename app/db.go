package app

import (
	"database/sql"
	"time"

	"github.com/jalal-akbar/belajar-golang-dependency-injection/helper"
)

func NewDB() *sql.DB {
	var (
		driverName     = "mysql"
		dataSourceName = "root:root@tcp(localhost:3306)/belajar_mysql_lagi"
	)
	db, err := sql.Open(driverName, dataSourceName)
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	return db
}
