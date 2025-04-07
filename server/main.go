package main

import (
	"fmt"
)

func main() {
	db := dbInit()
	if db == nil {
		fmt.Printf("db is status : %v", db)
	}
	server := newServerInit()

	router := server.newRouter("/cb")
	router("/test", test).TEST()
	router("/data", getData).GET()
	router("/chat", message).POST()
	router("/train", trainBot).POST()

	// enhancedRouter := enableCORS(jsonContentTypeMiddleware(router))

	server.run(":8080")

}
