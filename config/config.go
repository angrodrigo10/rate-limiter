package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	RateLimitIP    int
	RateLimitToken int
	BlockTime      int
}

func LoadConfig() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	ipLimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	tokenLimit, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	blockTime, _ := strconv.Atoi(os.Getenv("BLOCK_TIME"))

	return &Config{
		RedisAddr:      os.Getenv("REDIS_ADDR"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		RedisDB:        db,
		RateLimitIP:    ipLimit,
		RateLimitToken: tokenLimit,
		BlockTime:      blockTime,
	}
}
