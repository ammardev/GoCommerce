package main

import (
	"github.com/ammardev/gocommerce/internal/connections"
	"github.com/ammardev/gocommerce/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	connections.NewMySqlConnection()
	defer connections.CleanUp()

	server.Start()
}
