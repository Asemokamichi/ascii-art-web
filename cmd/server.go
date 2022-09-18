package cmd

import (
	"ascii-art/internal/delivery"
	"fmt"
	"log"
	"net/http"
)

func Servers() {
	server := delivery.New()
	fmt.Printf("Starting server at port 8080\nhttp://localhost:8080/\n")

	if err := http.ListenAndServe(":8080", server.Route()); err != nil {
		log.Fatal(err)
	}
}
