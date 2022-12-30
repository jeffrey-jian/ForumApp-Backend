package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CVWO/sample-go-app/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Println("Listening on port 8000 at http://localhost:8000!")

	log.Fatalln(http.ListenAndServe(":8000", r))
}
