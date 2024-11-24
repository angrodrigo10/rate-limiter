package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimiter struct {
	RedisClient *redis.Client
	IPLimit     int
	TokenLimit  int
	BlockTime   time.Duration
}

func NewRateLimiter(client *redis.Client, ipLimit, tokenLimit int, blockTime time.Duration) *RateLimiter {
	return &RateLimiter{
		RedisClient: client,
		IPLimit:     ipLimit,
		TokenLimit:  tokenLimit,
		BlockTime:   blockTime,
	}
}

func (rl *RateLimiter) Allow(ctx context.Context, key string, limit int) bool {
	val, err := rl.RedisClient.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return false
	}

	if val >= limit {
		rl.RedisClient.Set(ctx, key, val, rl.BlockTime).Err()
		return false
	}

	pipe := rl.RedisClient.Pipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Second)
	_, _ = pipe.Exec(ctx)

	return true
}
