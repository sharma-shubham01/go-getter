package main

import (
	"database/sql"
	"sync"
)

// sync.Map
// A concurrent map implementation optimized for specific cache-like patterns.
// Standard Go maps are not concurrency-safe; concurrent read/write causes a fatal runtime crash.

// Core API -
// Store(key, value any) / Load(key any) (value any, ok bool)
// LoadOrStore(key, value any) (actual any, loaded bool)
// Delete(key any) / LoadAndDelete(key any) (value any, loaded bool)
// Range(f func(key, value any) bool)

// Only reach for sync.Map when:
// Disjoint Key Sets: Multiple goroutines read, write, and overwrite disjoint sets of keys.
// Stable Keys (Write-Once, Read-Many): A key is written once and read millions of times (e.g., global configuration caches).

// sync.Once
// Ensures that a function runs exactly once across all goroutines, commonly used for lazy initialization and singletons.

var (
	once     sync.Once
	instance *sql.DB
)

func GetDB() *sql.DB {
	once.Do(func() {
		//instance = initDB() // Executed once; concurrent callers block until completed
	})
	return instance
}

// sync.Pool
// A pool of temporary objects that can be saved and reused to relieve Garbage Collection (GC) pressure.

// sync.Cond
// A condition variable used when goroutines need to wait for a state change or an event.
