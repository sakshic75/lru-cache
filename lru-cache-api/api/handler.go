package api

import (
	"encoding/json"
	"fmt"
	"lru-cache-api/utils"
	"net/http"
	"strconv"
	"time"
)

// Handler for setting a value in cache
func setHandler(w http.ResponseWriter, r *http.Request) {
	//enableCors(w)
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	durationStr := r.URL.Query().Get("duration")
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}

	utils.Cache.Set(key, value, time.Duration(duration)*time.Second)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": fmt.Sprintf("Success !! Set key: %s, value: %s, duration: %d seconds", key, value, duration)})

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	//enableCors(w)
	key := r.URL.Query().Get("key")
	value, found := utils.Cache.Get(key)
	if !found {
		http.Error(w, "Key not found or expired", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	//enableCors(w)
	key := r.URL.Query().Get("key")
	utils.Cache.Delete(key)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": fmt.Sprintf("Successfully Deleted key: %s", key)})

}
