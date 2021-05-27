package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPHost         string `envconfig:"HTTP_HOST"`
	HTTPPort         string `envconfig:"HTTP_PORT"`
	LogLevel         uint   `envconfig:"LOG_LEVEL"`
	PgURL            string `envconfig:"PG_URL"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
}

var (
	config Config
	once   sync.Once
)

//Get init Config, return *Config
func Get() *Config {
	once.Do(func() {
		//set envinment variables from file ".env"
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}

		//catch envinment variables to "&config"
		if err := envconfig.Process("", &config); err != nil {
			log.Fatal(err)
		}

		//catch flag loglevel or default from ".env"
		loglevel := flag.Uint("loglevel", config.LogLevel, "for logrus levels")

		flag.Parse() //parse flag

		config.LogLevel = *loglevel //override config.LogLevel

		configJsonBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Configuration: %v\n", string(configJsonBytes))
	})
	return &config
}
