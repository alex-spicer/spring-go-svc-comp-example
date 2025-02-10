package main

import (
	"log"
	"product-go/internal/server"
)

func main() {
	srv := server.New()
	log.Fatal(srv.Start(":8081"))
}
