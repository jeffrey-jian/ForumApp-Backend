package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/router"
)

func main() {
	port := "8000"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	fmt.Printf("Listening on port %s at http://localhost:%s\n", port, port)
	r := router.Setup()
	_, err := database.GetDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	log.Fatalln(http.ListenAndServe(":"+port, r))
}
