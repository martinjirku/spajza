package config

import "fmt"

func GetDBType() string {
	return DefaultConfiguration.DBType
}

func GetMariaDBSQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultConfiguration.DBUser,
		DefaultConfiguration.DBPassword,
		DefaultConfiguration.DBHost,
		DefaultConfiguration.DBPort,
		DefaultConfiguration.DBName)
	return dataBase
}
