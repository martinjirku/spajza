package main

import (
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

	userApp := users.NewUserApp(db)
	userApp.SetupRouter(e)

	unitsApp := units.NewUnitApp()
	unitsApp.SetupRouter(e)

	e.Logger.Fatal(web.StartWebServer(e))
}
