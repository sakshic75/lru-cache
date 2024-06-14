// utils/utils.go

package utils

import (
	"sync"
	"time"
)

// CacheItem represents a single item in the cache.
type CacheItem struct {
	Value      string
	Expiration int64
}

// LRUCache is a thread-safe fixed-size LRU cache.
type LRUCache struct {
	capacity int
	items    map[string]CacheItem
	mu       sync.Mutex
	order    []string
}

// NewLRUCache creates an LRUCache of the given capacity.
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		items:    make(map[string]CacheItem),
		order:    make([]string, 0, capacity),
	}
}

// Get retrieves a value from the cache.
func (c LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, found := c.items[key]; found {
		if time.Now().UnixNano() > item.Expiration {
			delete(c.items, key)
			c.removeFromOrder(key)
			return "", false
		}
		c.updateOrder(key)
		return item.Value, true
	}
	return "", false
}

// Set adds a value to the cache.
func (c LRUCache) Set(key, value string, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; found {
		c.updateOrder(key)
	} else {
		if len(c.items) >= c.capacity {
			oldest := c.order[0]
			c.order = c.order[1:]
			delete(c.items, oldest)
		}
		c.order = append(c.order, key)
	}

	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}

// Delete removes a value from the cache.
func (c LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; found {
		delete(c.items, key)
		c.removeFromOrder(key)
	}
}

// updateOrder moves a key to the end of the order slice.
func (c LRUCache) updateOrder(key string) {
	for i, k := range c.order {
		if k == key {
			c.order = append(c.order[:i], c.order[i+1:]...)
			c.order = append(c.order, key)
			break
		}
	}
}

// removeFromOrder removes a key from the order slice.
func (c LRUCache) removeFromOrder(key string) {
	for i, k := range c.order {
		if k == key {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
}

// Global cache instance
var Cache = NewLRUCache(5)
