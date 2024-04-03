package configs

import (
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("api.host", "localhost")
	viper.SetDefault("db.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = &Config{
		API: APIConfig{
			Host: viper.GetString("api.host"),
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			User:     viper.GetString("db.user"),
			Pass:     viper.GetString("db.pass"),
			Database: viper.GetString("db.name"),
		},
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
