package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvVariables()

	routes := router.CreateRouter()

	fmt.Printf("Server running on port %d\n", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), routes))
}
