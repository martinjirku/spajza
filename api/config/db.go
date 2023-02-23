package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetMariaDBSQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString(dbType),
		viper.GetString(dbPassword),
		viper.GetString(dbHost),
		viper.GetString(dbPort),
		viper.GetString(dbName))
	return dataBase
}
