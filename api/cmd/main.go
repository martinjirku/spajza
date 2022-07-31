package main

import (
	"github.com/martinjirku/zasobar/web"
)

func main() {
	e, err := web.CreateWebServer("8080")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(web.StartWebServer(e))
}
