package main

import (
	"fmt"
	"log"
	"net/http"

	"deniableEncryption/routes"
)

func main() {
	port := "3000"

	router := routes.SetupRoutes()

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
