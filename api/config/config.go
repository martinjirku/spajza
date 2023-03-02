package config

import (
	"github.com/spf13/viper"
)

type Jwt struct {
	Secret   string `mapstructure:"SECRET"`
	Validity int64  `mapstructure:"VALIDITY"`
	Issuer   string `mapstructure:"ISSUER"`
}

type Db struct {
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Type     string `mapstructure:"TYPE"`
}

type Configuration struct {
	viper  *viper.Viper
	Port   string `mapstructure:"PORT"`
	Domain string `mapstructure:"DOMAIN"`
	Jwt    Jwt    `mapstructure:"JWT"`
	DB     Db     `mapstructure:"DB"`
}

func NewConfiguration(v *viper.Viper) Configuration {
	return Configuration{viper: v}
}

func (c *Configuration) LoadConfiguration() error {
	return c.viper.Unmarshal(c)
}

const (
	portKey    = "port"
	domainKey  = "domain"
	jwtKey     = "jwt"
	dbKey      = "db"
	dbUser     = "db.user"
	dbPassword = "db.password"
	dbName     = "db.name"
	dbHost     = "db.host"
	dbPort     = "db.port"
	dbType     = "db.type"
)

func PrepareDefaultServe(v *viper.Viper, domainParam, portParam string) {
	setDefaultAndBindEnv(v, domainKey, domainParam)
	setDefaultAndBindEnv(v, portKey, portParam)
}

func PrepareJwt(v *viper.Viper) {
	jwt := Jwt{Secret: "Secret", Issuer: "Issuer", Validity: 10 * 60 * 24}
	setDefaultAndBindEnv(v, jwtKey, jwt)
}

func PrepareDefaults(v *viper.Viper) {
	db := Db{
		User:     "user",
		Password: "user",
		Name:     "zasobar",
		Host:     "localhost",
		Port:     "8000",
		Type:     "mysql",
	}
	setDefaultAndBindEnv(v, dbKey, db)
}

func setDefaultAndBindEnv(v *viper.Viper, key string, defaultValue interface{}) {
	v.BindEnv(key)
	v.SetDefault(key, defaultValue)
}
