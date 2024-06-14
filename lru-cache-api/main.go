// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"sync"
// 	"time"
// )

// // CacheItem represents a single item in the cache.
// type CacheItem struct {
// 	Value      string
// 	Expiration int64
// }

// // LRUCache is a thread-safe fixed-size LRU cache.
// type LRUCache struct {
// 	capacity int
// 	items    map[string]CacheItem
// 	mu       sync.Mutex
// 	order    []string
// }

// // NewLRUCache creates an LRUCache of the given capacity.
// func NewLRUCache(capacity int) LRUCache {
// 	return LRUCache{
// 		capacity: capacity,
// 		items:    make(map[string]CacheItem),
// 		order:    make([]string, 0, capacity),
// 	}
// }

// // Get retrieves a value from the cache.
// func (c LRUCache) Get(key string) (string, bool) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if item, found := c.items[key]; found {
// 		if time.Now().UnixNano() > item.Expiration {
// 			delete(c.items, key)
// 			c.removeFromOrder(key)
// 			return "", false
// 		}
// 		c.updateOrder(key)
// 		return item.Value, true
// 	}
// 	return "", false
// }

// // Set adds a value to the cache.
// func (c LRUCache) Set(key, value string, duration time.Duration) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if _, found := c.items[key]; found {
// 		c.updateOrder(key)
// 	} else {
// 		if len(c.items) >= c.capacity {
// 			oldest := c.order[0]
// 			c.order = c.order[1:]
// 			delete(c.items, oldest)
// 		}
// 		c.order = append(c.order, key)
// 	}

// 	c.items[key] = CacheItem{
// 		Value:      value,
// 		Expiration: time.Now().Add(duration).UnixNano(),
// 	}
// }

// // Delete removes a value from the cache.
// func (c LRUCache) Delete(key string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	if _, found := c.items[key]; found {
// 		delete(c.items, key)
// 		c.removeFromOrder(key)
// 	}
// }

// // updateOrder moves a key to the end of the order slice.
// func (c LRUCache) updateOrder(key string) {
// 	for i, k := range c.order {
// 		if k == key {
// 			c.order = append(c.order[:i], c.order[i+1:]...)
// 			c.order = append(c.order, key)
// 			break
// 		}
// 	}
// }

// // removeFromOrder removes a key from the order slice.
// func (c LRUCache) removeFromOrder(key string) {
// 	for i, k := range c.order {
// 		if k == key {
// 			c.order = append(c.order[:i], c.order[i+1:]...)
// 			break
// 		}
// 	}
// }

// // Global cache instance
// var cache = NewLRUCache(5)

// func setHandler(w http.ResponseWriter, r *http.Request) {
// 	enableCors(w)
// 	key := r.URL.Query().Get("key")
// 	value := r.URL.Query().Get("value")
// 	durationStr := r.URL.Query().Get("duration")
// 	duration, err := strconv.Atoi(durationStr)
// 	if err != nil {
// 		http.Error(w, "Invalid duration", http.StatusBadRequest)
// 		return
// 	}

// 	cache.Set(key, value, time.Duration(duration)*time.Second)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("Set key: %s, value: %s, duration: %d seconds", key, value, duration)))
// }

// func getHandler(w http.ResponseWriter, r *http.Request) {
// 	enableCors(w)
// 	key := r.URL.Query().Get("key")
// 	value, found := cache.Get(key)
// 	if !found {
// 		http.Error(w, "Key not found or expired", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{"value": value})
// }

// func deleteHandler(w http.ResponseWriter, r *http.Request) {
// 	enableCors(w)
// 	key := r.URL.Query().Get("key")
// 	cache.Delete(key)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("Deleted key: %s", key)))
// }

// func enableCors(w http.ResponseWriter) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

// func main() {
// 	http.HandleFunc("/set", setHandler)
// 	http.HandleFunc("/get", getHandler)
// 	http.HandleFunc("/delete", deleteHandler)

// 	log.Println("Server started on :8005")
// 	log.Fatal(http.ListenAndServe(":8005", nil))
// }

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"lru-cache-api/api"
	// Using gorilla/mux for routing
)

func main() {
	router := api.InitializeRouter()

	// CORS configuration
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Adjust to your frontend URL
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})

	// CORS middleware
	corsHandler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	log.Println("Server started on :8005")
	log.Fatal(http.ListenAndServe(":8005", corsHandler))
}
