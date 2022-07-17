package main

import (
	"github.com/martinjirku/zasobar/categories"
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
	db := storage.NewDB()

	users.StartApp(db, e)
	units.StartApp(e)
	categories.StartApp(db, e)

	e.Logger.Fatal(web.StartWebServer(e))
}
