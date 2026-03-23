package config

import (
	"os"
)

type Config struct {
	DBUrl string
}

func Load() *Config {
	dbUrl := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") + "?sslmode=disable"

	return &Config{
		DBUrl: dbUrl,
	}
}
