package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	api := APIConfig{
		Host: os.Getenv("API_HOST"),
		Port: os.Getenv("API_PORT"),
	}

	db := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
	redisDb, _ := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 32)
	cache := RedisCache{
		Host: os.Getenv("REDIS_HOST"),
		Db:   redisDb,
	}

	return &Config{
		API:   api,
		DB:    db,
		Cache: cache,
	}, err
}

type Config struct {
	API   APIConfig
	DB    DBConfig
	Cache RedisCache
}

type APIConfig struct {
	Host string
	Port string
}
