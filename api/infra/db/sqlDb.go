package db

import (
	"database/sql"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	config "github.com/martinjirku/zasobar/config"
	"github.com/spf13/viper"
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

var SqlDb *sql.DB

func init() {
	mx := sync.Mutex{}
	SqlDb = NewDB()
	viper.GetViper().OnConfigChange(func(in fsnotify.Event) {
		mx.Lock()
		SqlDb = NewDB()
		mx.Unlock()
	})
}
