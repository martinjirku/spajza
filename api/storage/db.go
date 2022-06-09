package storage

import (
	"log"

	config "github.com/martinjirku/zasobar/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMariaDBSQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(mysql.New(mysql.Config{DSN: conString}))

	if err != nil {
		log.Panic(err)
	}

	return DB
}
