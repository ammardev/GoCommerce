package main

import (
	"github.com/ammardev/ecommerce-playground/internal/connections"
	"github.com/ammardev/ecommerce-playground/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	connections.NewMySqlConnection()
	defer connections.CleanUp()

	server.Start()
}
