package main

import "sync"

// Mutex

// sync.Mutex provides mutual exclusion to serialize access to shared memory.

// Core Mechanics -
// Lock(): Acquires the lock. If held by another goroutine, the caller blocks.
// Unlock(): Releases the lock. Unlocking an unlocked mutex triggers a runtime panic.

type SafeCounter struct {
	mu sync.Mutex
	v  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock() // Idiomatic defer ensures release on panics/returns
	c.v++
}

// Important Points -
// Copying a Mutex: Passing a struct containing a Mutex by value copies its internal lock state, breaking synchronization.

// Deadlocks: Forgetting to unlock, re-locking in the same goroutine (Go mutexes are non-reentrant/recursive),
// or acquiring multiple locks in inconsistent order.

// Normal vs. Starvation Mode (Advanced):
// Normal mode: Goroutines compete in FIFO order, but newly woken goroutines have a CPU cache advantage, meaning a waiting goroutine can starve.
// Starvation mode: If a goroutine waits more than 1ms, the lock hands off ownership directly to the head of the wait queue to ensure fairness.

// --------------------------------------------------------------------------------------------------------------------------------------------------

// RWMutex

// sync.RWMutex
// A reader/writer mutual exclusion lock. Multiple readers can hold the lock simultaneously, or a single writer can hold it exclusively.

// Reader vs. Writer Scenarios
// Read Lock (RLock() / RUnlock()): Use when merely inspecting, serializing, or fetching data without mutating internal state.
// Write Lock (Lock() / Unlock()): Use when inserting, deleting, updating, or mutating the underlying data structure.

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

// Get Read scenario: Multiple goroutines can execute this simultaneously
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

// Set Write scenario: Blocks all current and pending reads until complete
func (c *Cache) Set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

// Important Points -
// A common pitfall is assuming RWMutex is always faster than Mutex.
// Under high-write contention, RWMutex performs worse due to additional book-keeping overhead.

// Go's RWMutex gives priority to incoming writers.
// When a writer calls Lock(), new readers are blocked from acquiring RLock(),
// allowing the writer to acquire the lock as soon as the existing readers finish.
