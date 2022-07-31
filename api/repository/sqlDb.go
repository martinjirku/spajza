package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/martinjirku/zasobar/config"
)

func NewDB() *sql.DB {
	var err error
	conString := config.GetMariaDBSQLConnectionString()

	DB, err := sql.Open("mysql", conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

var SqlDb = NewDB()
