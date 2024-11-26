package limiter

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter_Allow(t *testing.T) {
	// Inicializa um servidor Redis em memória usando o miniredis
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("não foi possível inicializar o miniredis: %v", err)
	}
	defer mr.Close()

	// Configura o cliente Redis para usar o servidor miniredis
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	// Contexto para os métodos do Redis
	ctx := context.Background()

	// Configura o RateLimiter
	ipLimit := 5
	blockTime := 10 * time.Second
	rateLimiter := NewRateLimiter(client, ipLimit, 0, blockTime)

	// Chave a ser usada no teste
	key := "test-ip"

	// Testa as requisições dentro do limite
	for i := 0; i < ipLimit; i++ {
		allowed := rateLimiter.Allow(ctx, key, ipLimit)
		assert.True(t, allowed, "A requisição deveria ser permitida")
	}

	// Testa uma requisição além do limite
	allowed := rateLimiter.Allow(ctx, key, ipLimit)
	assert.False(t, allowed, "A requisição deveria ser bloqueada")

	// Verifica se o bloqueio foi aplicado corretamente
	ttl := mr.TTL(key)
	assert.InEpsilon(t, blockTime.Seconds(), ttl.Seconds(), 1, "O tempo de bloqueio está incorreto")
}
