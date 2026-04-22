package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Application struct {
		Name    string `yaml:"name" env:"APP_NAME" env-default:"Grind"`
		Startup struct {
			Script string `yaml:"script" env:"STARTUP_SCRIPT" env-default:"startup.lua"`
		} `yaml:"startup"`
		Screen struct {
			Width      int    `yaml:"width" env:"SCREEN_WIDTH" env-default:"1024"`
			Height     int    `yaml:"height" env:"SCREEN_HEIGHT" env-default:"768"`
			Background string `yaml:"background" env:"SCREEN_BACKGROUND" env-default:"background.png"`
		} `yaml:"screen"`
		Hero struct {
			Atlas   string `yaml:"atlas" env:"HERO_ATLAS" env-default:"runner.png"`
			Scale   int    `yaml:"scale" env:"HERO_SCALE" env-default:"64"`
			Unit    int    `yaml:"unit" env:"HERO_UNIT" env-default:"16"`
			GroundY int    `yaml:"ground" env:"HERO_GROUND_Y" env-default:"680"`
		} `yaml:"hero"`
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
