package configs

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort         string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPass          string
	DBName          string
	DBSSLMode       string
	RedisHost       string
	RedisPort       string
	RedisPassword   string
	RedisDB         string
	VKMiniAppID     string
	VKMiniAppSecret string
	VKMiniAppToken  string
}

func GetConfig() (Config, error) {

	AppPort := os.Getenv("APP_PORT")
	if AppPort == "" {
		AppPort = "8080"
	}
	DBHost := os.Getenv("DB_HOST")
	if DBHost == "" {
		return Config{}, fmt.Errorf("DB_HOST is not set")
	}
	DBPort := os.Getenv("DB_PORT")
	if DBPort == "" {
		return Config{}, fmt.Errorf("DB_PORT is not set")
	}
	DBUser := os.Getenv("DB_USER")
	if DBUser == "" {
		return Config{}, fmt.Errorf("DB_USER is not set")
	}
	DBPass := os.Getenv("DB_PASS")
	if DBPass == "" {
		return Config{}, fmt.Errorf("DB_PASS is not set")
	}
	DBName := os.Getenv("DB_NAME")
	if DBName == "" {
		return Config{}, fmt.Errorf("DB_NAME is not set")
	}
	DBSSLMode := os.Getenv("DB_SSL_MODE")
	if DBSSLMode == "" {
		return Config{}, fmt.Errorf("DB_SSL_MODE is not set")
	}

	RedisHost := os.Getenv("REDIS_HOST")
	if RedisHost == "" {
		return Config{}, fmt.Errorf("REDIS_HOST is not set")
	}
	RedisPort := os.Getenv("REDIS_PORT")
	if RedisPort == "" {
		return Config{}, fmt.Errorf("REDIS_PORT is not set")
	}
	RedisPassword := os.Getenv("REDIS_PASSWORD")
	if RedisPassword == "" {
		return Config{}, fmt.Errorf("REDIS_PASSWORD is not set")
	}
	RedisDB := os.Getenv("REDIS_DB")
	if RedisDB == "" {
		return Config{}, fmt.Errorf("REDIS_DB is not set")
	}

	cfg := Config{
		AppPort:       AppPort,
		DBHost:        DBHost,
		DBPort:        DBPort,
		DBUser:        DBUser,
		DBPass:        DBPass,
		DBName:        DBName,
		DBSSLMode:     DBSSLMode,
		RedisHost:     RedisHost,
		RedisPort:     RedisPort,
		RedisPassword: RedisPassword,
		RedisDB:       RedisDB,
	}

	return cfg, nil
}
