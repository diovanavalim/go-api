package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes := router.CreateRouter()

	port := ":8686"

	fmt.Printf("Server running on port %s", port)

	log.Fatal(http.ListenAndServe(port, routes))
}
