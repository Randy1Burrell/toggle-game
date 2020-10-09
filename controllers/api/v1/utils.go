package v1

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/randy1burrell/toggle-game/pkg/logger"
)

// Handler interface for httpHandlers
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// NotFound handler
type NotFound struct {
	count int
	mu    sync.Mutex
}

// ServeHTTP for not found
func (c *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%d) Requested path %s not found", c.count, r.URL.Path)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(`{"message": "Path not found"}`)
	return
}

// NotAllowed handler
type NotAllowed struct {
	count int
	mu    sync.Mutex
}

// ServeHTTP for not allowed methods
func (c *NotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Printf("%d) Requested method %s not allowed for path %s", c.count, r.Method, r.URL.Path)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(`{"message": "Method not allowed"}`)
	return
}
