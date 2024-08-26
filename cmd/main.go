package main

import "github.com/OlyMahmudMugdho/ecommerce-api-gateway/server"

func main() {
	server := server.NewServer("8080")
	server.Run()
}
