package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var mu sync.Mutex
var clients = make(map[string]*clientLimiter)
var cleanupInterval = time.Minute * 5

func InitRateLimiter() {
	go clenupOldClients()
}

func clenupOldClients() {
	for {
		time.Sleep(cleanupInterval)

		mu.Lock()

		for ip, cl := range clients {
			if time.Since(cl.lastSeen) > cleanupInterval {
				delete(clients, ip)
			}
		}

		mu.Unlock()
	}
}

func getClientIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	cl, exists := clients[ip]

	if !exists {
		limiter := rate.NewLimiter(1, 5)
		clients[ip] = &clientLimiter{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
	}

	cl.lastSeen = time.Now()
	return cl.limiter
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getClientIP(r)
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "Too many request", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
