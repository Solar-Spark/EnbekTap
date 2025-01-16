package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type ClientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu              sync.Mutex
	clients         = make(map[string]*ClientLimiter)
	maxRate         = rate.Every(1 * time.Second)
	burst           = 5
	cleanupInterval = 5 * time.Minute
)

func RateLimiterMiddleware() gin.HandlerFunc {
	go cleanupClients()

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		limiter := getClientLimiter(clientIP)

		if !limiter.limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func getClientLimiter(clientIP string) *ClientLimiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, exists := clients[clientIP]; exists {
		limiter.lastSeen = time.Now()
		return limiter
	}

	limiter := &ClientLimiter{
		limiter:  rate.NewLimiter(maxRate, burst),
		lastSeen: time.Now(),
	}
	clients[clientIP] = limiter
	return limiter
}

func cleanupClients() {
	for {
		time.Sleep(cleanupInterval)

		mu.Lock()
		for clientIP, limiter := range clients {
			if time.Since(limiter.lastSeen) > cleanupInterval {
				delete(clients, clientIP)
			}
		}
		mu.Unlock()
	}
}
