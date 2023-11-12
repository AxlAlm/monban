package config

import (
	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	DBHost         string
	DBPort         int
	DBUsername     string
	DBPassword     string
	DBName         string
	DBMaxPoolConns int
}

func init() {
	viper.SetConfigType("env")
	viper.SetEnvPrefix("MONBAN")
	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USERNAME", "monban")
	viper.SetDefault("DB_PASSWORD", "monban")
	viper.SetDefault("DB_NAME", "monban")
	viper.SetDefault("DB_MAX_POOL_CONNS", 5)

	Config = Configuration{
		DBHost:         viper.GetString("DB_HOST"),
		DBPort:         viper.GetInt("DB_PORT"),
		DBUsername:     viper.GetString("DB_USERNAME"),
		DBPassword:     viper.GetString("DB_PASSWORD"),
		DBName:         viper.GetString("DB_NAME"),
		DBMaxPoolConns: viper.GetInt("DB_MAX_POOL_CONNS"),
	}

}
