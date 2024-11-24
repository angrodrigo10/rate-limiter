package server

import (
	"time"

	"github.com/angrodrigo10/rate-limiter/config"
	"github.com/angrodrigo10/rate-limiter/internal/limiter"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func StartServer(cfg *config.Config) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	rateLimiter := limiter.NewRateLimiter(rdb, cfg.RateLimitIP, cfg.RateLimitToken, time.Duration(cfg.BlockTime)*time.Second)

	r := gin.Default()
	r.Use(limiter.RateLimitMiddleware(rateLimiter))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome!",
		})
	})

	r.Run(":8080")
}
