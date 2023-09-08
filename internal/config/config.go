package config

import (
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "cestevezing"),
		Password: getEnv("DB_PASSWORD", "mysecret"),
		Name:     getEnv("DB_NAME", "eventos"),
	}
}

func getEnv(key, optionalValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return optionalValue
}
