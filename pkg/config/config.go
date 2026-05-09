package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Application struct {
		Name string `yaml:"name" env:"APP_NAME" env-default:"Grind"`
		Log  struct {
			Level int `yaml:"level" env:"LOG_LEVEL" env-default:"0"`
		} `yaml:"log"`
	} `yaml:"application"`
}

func Load(configDir string) *Config {
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "production"
	}

	var configuration Config

	configFile := fmt.Sprintf("%s/config.%s.yaml", configDir, environment)
	log.Printf("Reading configuration file %s", configFile)
	if err := cleanenv.ReadConfig(configFile, &configuration); err != nil {
		log.Printf("Failed to read configuration %s", configFile)
		cleanenv.ReadEnv(&configuration)
	}

	localFile := fmt.Sprintf("%s/config.local.yaml", configDir)
	log.Printf("Reading local configuration file %s", localFile)
	if _, err := os.Stat(localFile); err == nil {
		if err := cleanenv.ReadConfig(localFile, &configuration); err != nil {
			log.Printf("Error reading configuration %s: %v", localFile, err)
		}
	}

	return &configuration
}
