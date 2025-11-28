package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port string

	DBUser string
	DBPassword string
	DBName string
	DBHost string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config {
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName: getEnv("DB_NAME", "db"),
		DBHost: getEnv("DB_HOST", "db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}