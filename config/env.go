package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host          string
	Port          string
	DBUser        string
	DBPassword    string
	DBAddress     string
	DBName        string
	JWTSecret     string
	JWTExpiration int
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:          getStrEnv("HOST", "http://localhost:"),
		Port:          getStrEnv("PORT", "4444"),
		DBUser:        getStrEnv("DB_USER", "root"),
		DBPassword:    getStrEnv("DB_PASSWORD", "root1"),
		DBAddress:     getStrEnv("DB_ADDRESS", "127.0.0.1"),
		DBName:        getStrEnv("DB_NAME", "password-service"),
		JWTSecret:     getStrEnv("JWT_SECRET", "secret_Key"),
		JWTExpiration: getIntEnv("JWT_EXPIRATION", 3600),
	}
}

func getStrEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}
	return fallback
}
