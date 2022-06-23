package model

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Config() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tempone_db")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Banco de dados conectado!")
}
