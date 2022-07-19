package config

import (
	"sync"

	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env              string `mapstructure:"ENV"`
	Port             string `mapstructure:"PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresSchema   string `mapstructure:"POSTGRES_SCHEMA"`
	PostgresUsername string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
}

var config *Config
var once sync.Once

func init() {
	once.Do(func() {
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")
		config = new(Config)
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, %s", err)

		}
		if err := viper.Unmarshal(config); err != nil {
			log.Printf("Unable to decode into struct, %v", err)

		}

	})
}

func GetConfig() *Config {
	return config
}
