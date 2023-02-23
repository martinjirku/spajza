package config

import (
	"github.com/spf13/viper"
)

type Jwt struct {
	Secret   string
	Validity uint
	Issuer   string
}

type Db struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
	Type     string
}

type Configuration struct {
	Port   string
	Domain string
	Jwt    Jwt
	DB     Db
}

const (
	port        = "port"
	domain      = "domain"
	jwtSecret   = "jwt.secret"
	jwtValidity = "Jwt.Validity"
	dbUser      = "db.user"
	dbPassword  = "db.password"
	dbName      = "db.name"
	dbHost      = "db.host"
	dbPort      = "db.port"
	dbType      = "db.type"
)

func GetJwtSecret() string {
	return viper.GetString(jwtSecret)
}

func PrepareDefaults() {
	viper.SetDefault(domain, "localhost")
	viper.SetDefault(port, "8000")
	viper.SetDefault(jwtSecret, "TopSecret")
	viper.SetDefault(jwtValidity, 10*60*24)
	viper.SetDefault(dbUser, "user")
	viper.SetDefault(dbPassword, "user")
	viper.SetDefault(dbName, "zasobar")
	viper.SetDefault(dbHost, "localhost")
	viper.SetDefault(dbPort, "3306")
	viper.SetDefault(dbType, "mysql")
}

func SetConfigurations() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/zasobar/")
	viper.AddConfigPath(".")
}

func GetConfiguration() Configuration {
	config := Configuration{}
	viper.Unmarshal(&config)
	return config
}
