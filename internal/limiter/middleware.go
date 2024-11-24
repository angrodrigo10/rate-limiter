package limiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		token := c.GetHeader("API_KEY")

		key := ip
		limit := rl.IPLimit

		if token != "" {
			key = "token:" + token
			limit = rl.TokenLimit
		}

		if !rl.Allow(c.Request.Context(), key, limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "you have reached the maximum number of requests or actions allowed within a certain time frame",
			})
			return
		}

		c.Next()
	}
}
