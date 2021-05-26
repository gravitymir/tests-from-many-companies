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
	HTTPAddr         string `envconfig:"HTTP_ADDR"`
	PgURL            string `envconfig:"PG_URL"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	LogLevel         uint   `envconfig:"LOG_LEVEL"`
}

var (
	config Config
	once   sync.Once
)

//Get init once Config, return always *Config
func Get() *Config {

	once.Do(func() {

		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}

		if err := envconfig.Process("", &config); err != nil {
			log.Fatal(err)
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		//catch flag loglevel
		loglevel := flag.Uint("loglevel", config.LogLevel, "for logrus levels")

		flag.Parse()

		config.LogLevel = *loglevel

		fmt.Println("Configuration:", string(configBytes))
	})

	return &config
}
