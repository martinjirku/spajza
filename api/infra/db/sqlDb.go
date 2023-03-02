package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/martinjirku/zasobar/config"
)

func NewDB(dbConfig config.Db) *sql.DB {
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	DB, err := sql.Open("mysql", conString)
	if err != nil {
		log.Panic(err)
	}

	return DB
}
