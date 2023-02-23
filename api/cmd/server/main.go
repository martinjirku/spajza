package main

import (
	"fmt"

	"github.com/martinjirku/zasobar/config"
	web "github.com/martinjirku/zasobar/infra/web"
	"github.com/spf13/viper"
)

func main() {
	config.PrepareDefaults()
	config.SetConfigurations()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	web.InitServer()
}
