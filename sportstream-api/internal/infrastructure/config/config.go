package config

import (
	"os"
	"strings"
)

type Config struct {
	Port               string
	Env                string
	Version            string
	CORSAllowedOrigins []string
}

func Load() *Config {
	return &Config{
		Port:               getEnv("APP_PORT", "8080"),
		Env:                getEnv("APP_ENV", "development"),
		Version:            getEnv("APP_VERSION", "1.0.0"),
		CORSAllowedOrigins: strings.Split(getEnv("CORS_ALLOWED_ORIGINS", "*"), ","),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
