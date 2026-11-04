package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RateLimiter holds the state for token bucket algorithm
type RateLimiter struct {
	mu        sync.Mutex
	tokens    int           // Current available tokens
	maxTokens int           // Maximum capacity of the bucket
	refillHz  time.Duration // Frequency at which 1 token is added back
}

// NewRateLimiter initializes the limiter and starts the background refill routine
func NewRateLimiter(maxTokens int, refillHz time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:    maxTokens,
		maxTokens: maxTokens,
		refillHz:  refillHz,
	}

	// Background ticker goroutine to replenish tokens continuously
	go func() {
		ticker := time.NewTicker(rl.refillHz)
		defer ticker.Stop()

		for range ticker.C {
			rl.mu.Lock()
			if rl.tokens < rl.maxTokens {
				rl.tokens++
			}
			rl.mu.Unlock()
		}
	}()

	return rl
}

// Middleware returns an standard HTTP middleware wrapper
func (rl *RateLimiter) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()

		// 1. Check if tokens are available
		if rl.tokens <= 0 {
			rl.mu.Unlock()
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// 2. Consume 1 token and allow request through
		rl.tokens--
		rl.mu.Unlock()

		next.ServeHTTP(w, r)
	}
}

// Dummy Endpoint for testing
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request successful!"))
}

func main() {
	// Allow a max burst of 5 requests, refilling 1 token every 500ms
	limiter := NewRateLimiter(5, 500*time.Millisecond)

	http.HandleFunc("/api", limiter.Middleware(helloHandler))

	fmt.Println("Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}
