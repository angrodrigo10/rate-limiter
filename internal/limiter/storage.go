package limiter

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisStorage é uma implementação do StorageInterface usando Redis.
type RedisStorage struct {
	Client *redis.Client
}

// NewRedisStorage cria uma instância do RedisStorage.
func NewRedisStorage(client *redis.Client) *RedisStorage {
	return &RedisStorage{Client: client}
}

// Increment incrementa o valor de um contador em Redis.
func (r *RedisStorage) Increment(ctx context.Context, key string) (int, error) {
	val, err := r.Client.Incr(ctx, key).Result() // Result retorna o valor incrementado.
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

// Get obtém o valor associado a uma chave no Redis.
func (r *RedisStorage) Get(ctx context.Context, key string) (int, error) {
	val, err := r.Client.Get(ctx, key).Result() // Result retorna o valor como string.
	if err != nil {
		if err == redis.Nil {
			// Chave não encontrada.
			return 0, nil
		}
		return 0, err
	}

	parsedVal, parseErr := strconv.Atoi(val)
	if parseErr != nil {
		return 0, parseErr
	}

	return parsedVal, nil
}

// SetExpiry define o tempo de expiração para uma chave no Redis.
func (r *RedisStorage) SetExpiry(ctx context.Context, key string, expiry time.Duration) error {
	_, err := r.Client.Expire(ctx, key, expiry).Result() // Result retorna true se a expiração foi configurada.
	return err
}
