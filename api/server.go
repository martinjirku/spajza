package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/martinjirku/zasobar/categories"
	"github.com/martinjirku/zasobar/config"
	"github.com/martinjirku/zasobar/units"
	"github.com/martinjirku/zasobar/users"
	"github.com/martinjirku/zasobar/web"

	"github.com/martinjirku/zasobar/storage"
)

func main() {
	e, err := web.CreateWebServer("8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
	gormDb := storage.NewDB()

	db, err := sql.Open("mysql", config.GetMariaDBSQLConnectionString())
	if err != nil {
		e.Logger.Fatal("Failed createDb")
	}

	users.StartApp(gormDb, db, e)
	units.StartApp(e)
	categories.StartApp(db, e)

	e.Logger.Fatal(web.StartWebServer(e))
}
