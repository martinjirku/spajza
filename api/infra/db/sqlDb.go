package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	maskedString := strings.ReplaceAll(conString, dbConfig.Password, "***")
	log.Default().Printf("connection string %q created", maskedString)

	DB, err := sql.Open("mysql", conString)
	if err != nil {
		log.Panicf("Could not open an database %q. Failed at %s", maskedString, err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Could not connect to db: %q", err)
	}
	log.Default().Printf("db connected %q created", maskedString)
	return DB
}
