package config

type Configuration struct {
	JWTSecret   string
	JWTValidity uint
	JWTIssuer   string
	DBUser      string
	DBPassword  string
	DBName      string
	DBHost      string
	DBPort      string
	DBType      string
}

var DefaultConfiguration = &Configuration{
	JWTSecret:   "TopSecret",
	JWTValidity: 10,
	JWTIssuer:   "zasobar",
	DBUser:      "user",
	DBPassword:  "user",
	DBName:      "zasobar",
	DBHost:      "localhost",
	DBPort:      "3306",
	DBType:      "mysql",
}

func GetJwtSecret() string {
	return DefaultConfiguration.JWTSecret
}
