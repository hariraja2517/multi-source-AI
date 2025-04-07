package main

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

func bindJson(r *http.Request, v any) {
	json.NewDecoder(r.Body).Decode(v)
}

// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Set CORS headers
// 		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		// Check if the request is for CORS preflight
// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		// Pass down the request to the next middleware (or final handler)
// 		next.ServeHTTP(w, r)
// 	})

// }

// func jsonContentTypeMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Set JSON Content-Type
// 		w.Header().Set("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	})
// }
