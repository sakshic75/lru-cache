package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() http.Handler {
	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/set", setHandler).Methods("GET")
	r.HandleFunc("/get", getHandler).Methods("GET")
	r.HandleFunc("/delete", deleteHandler).Methods("DELETE")
	r.HandleFunc("/options", deleteHandler).Methods("OPTIONS")

	return r
}
