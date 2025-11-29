package config

import (
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port string

	DBUser string
	DBPassword string
	DBName string
	DBHost string
	JWTSecret string
	JWTExpirationInSeconds int64
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
		JWTSecret: getEnv("JWT_SECRET", "not-secret-anymore"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRE_TIME", 900),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallbalck int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallbalck
		}

		return i
	}

	return fallbalck
}