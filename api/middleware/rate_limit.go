package middleware

import (
	"net/http"
	"sync"

	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type ClientLimiter struct {
	clients map[string]*rate.Limiter
	mu      sync.Mutex
	r       rate.Limit
	burst   int
}

func NewClientLimiter(r rate.Limit, b int) *ClientLimiter {
	return &ClientLimiter{
		clients: make(map[string]*rate.Limiter),
		r:       r,
		burst:   b,
	}
}

func (cl *ClientLimiter) getLimiter(key string) *rate.Limiter {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	limiter, exists := cl.clients[key]
	if !exists {
		limiter = rate.NewLimiter(cl.r, cl.burst)
		cl.clients[key] = limiter
	}
	return limiter
}

func RateLimitMiddleware(cl *ClientLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var key string

		if val, exists := c.Get("userID"); exists {
			key = fmt.Sprintf("user_%v", val)
		} else {
			key = fmt.Sprintf("ip_%s", c.ClientIP())
		}
		limiter := cl.getLimiter(key)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, try again later.",
			})
			return
		}
		c.Next()
	}
}
