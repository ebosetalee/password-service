package config

import (
	"os"

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
	JWTExpiration int64
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Host:       getStrEnv("HOST", "http://localhost:"),
		Port:       getStrEnv("PORT", "4444"),
		DBUser:     getStrEnv("DB_USER", "root"),
		DBPassword: getStrEnv("DB_PASSWORD", "root1"),
		DBAddress:  getStrEnv("DB_ADDRESS", "127.0.0.1"),
		DBName:     getStrEnv("DB_NAME", "password-service"),
	}
}

func getStrEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// func getIntEnv(key, fallback int64) int64 {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}
// 	return fallback
// }
