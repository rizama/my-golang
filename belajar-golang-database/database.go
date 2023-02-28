package belajar_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:sam123@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                 // minimal connection pertama kali dibuat untuk standby
	db.SetMaxOpenConns(100)                // maximal conenction yang akan dibuat oleh golang
	db.SetConnMaxIdleTime(time.Minute * 1) // berapa lama connection baru (diluar jumlah minimum yaitu 10) akan dihapus ketika sudah tidak dipakai
	db.SetConnMaxLifetime(time.Hour * 1)   // berapa lama connection akan di renew / diperbaharui, termasuk connection minimum yang standby (10 connection)

	return db
}
