package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CVWO/sample-go-app/internal/router"
)

func main() {
	port := "8000"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	fmt.Printf("Listening on port %s at http://localhost:%s\n", port, port)
	r := router.Setup()
	log.Fatalln(http.ListenAndServe(":"+port, r))
}
