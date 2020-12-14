package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	API struct {
		Host string `envconfig:"HOST" default:":3131"`
	}
	DB struct {
		Host string `envconfig:"HOST" required:"true"`
		Port string `envconfig:"PORT" required:"true"`
		User string `envconfig:"USER" required:"true"`
		Pass string `envconfig:"PASS" required:"true"`
		Name string `envconfig:"NAME" required:"true"`
	}
}

func loadEnv(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

func LoadConfig() (*Configuration, error) {
	if err := loadEnv("./.env"); err != nil {
		return nil, err
	}

	config := new(Configuration)
	if err := envconfig.Process("APP", config); err != nil {
		return nil, err
	}
	return config, nil
}
