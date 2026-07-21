package settings

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string `env:"DB_URL"`
	DBPort     string `env:"DB_PORT"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`

	ServerPort string `env:"SERVER_PORT" envDefault:"3030"`
	ServerURI  string `env:"SERVER_URI" envDefault:"http://localhost"`
}

func NewConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("Variáveis do .env não carregadas")
	}

	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
