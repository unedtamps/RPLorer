package middleware

import (
	"errors"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unedtamps/go-backend/config"
	"github.com/unedtamps/go-backend/util"
	"golang.org/x/time/rate"
)

type client struct {
	limit     *rate.Limiter
	last_seen time.Time
}

// rw mutex use for read and write in concurent
var (
	mutex   sync.RWMutex
	clients = make(map[string]*client)
)

func RateLimit(c *gin.Context) {

	go clearClient()

	if config.Env.Enable {
		ip := c.RemoteIP()
		mutex.Lock()

		defer mutex.Unlock()

		if _, ok := clients[ip]; !ok {
			clients[ip] = &client{
				limit: rate.NewLimiter(rate.Limit(config.Env.Rps), config.Env.Burst),
			}
		}
		clients[ip].last_seen = time.Now()
		if !clients[ip].limit.Allow() {
			util.LimitError(c, errors.New("Rate Limit Exceeded"))
			c.Abort()
			return
		}
	}
	c.Next()
}

func clearClient() {
	for {
		time.Sleep(1 * time.Minute)
		mutex.Lock()
		for ip, v := range clients {
			if time.Since(v.last_seen) > 1*time.Hour {
				delete(clients, ip)
			}
		}
		mutex.Unlock()
	}

}
