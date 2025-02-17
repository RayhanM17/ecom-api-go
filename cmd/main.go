package main

import (
	"log"

	"github.com/rayhanm17/ecom-api-go/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
