package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load variables
	config.Load()

	// Create all routers
	r := router.Create()

	// Running the server
	fmt.Printf("Listen on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}