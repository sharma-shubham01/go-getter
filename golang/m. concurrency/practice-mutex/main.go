package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Implement a thread-safe round-robin load balancer in Go.
// You are provided with a list of backend IP addresses.
// Write a method that returns the next IP address for each request, looping back to the beginning when the list ends.
// Ensure your implementation is safe for concurrent use when multiple goroutines call the method simultaneously.
// Do not use any external packages beyond the Go standard library.

type LoadBalancer struct {
	addresses []string
	index     int
	mu        sync.RWMutex

	// for atomic
	counter atomic.Uint64
}

// using mutex
func (lb *LoadBalancer) Next() (string, error) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	if len(lb.addresses) == 0 {
		return "", fmt.Errorf("no IP addresses found")
	}

	ip := lb.addresses[lb.index]
	lb.index = (lb.index + 1) % len(lb.addresses)

	return ip, nil
}

// using sync/atomic
func (lb *LoadBalancer) NextUsingAtomic() (string, error) {
	if len(lb.addresses) == 0 {
		return "", nil
	}

	n := lb.counter.Add(1) - 1

	lb.index = int(n % uint64(len(lb.addresses)))

	return lb.addresses[lb.index], nil
}

func main() {
	lb := LoadBalancer{
		addresses: []string{
			"192.168.1.1",
			"192.168.1.2",
			"192.168.1.3",
		},
	}

	var wg sync.WaitGroup

	for i := 1; i <= 7; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			//ip, err := lb.Next()
			ip, err := lb.NextUsingAtomic()
			if err != nil {
				fmt.Printf("Error occurred %v\n", err)
				return
			}

			fmt.Printf("Go Routine: %d, IP Address: %s\n", id, ip)
		}(i)
	}

	wg.Wait()
}
