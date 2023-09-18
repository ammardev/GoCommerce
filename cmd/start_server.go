package main

import (
	"github.com/ammardev/gocommerce/internal/connections"
	"github.com/ammardev/gocommerce/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	defer connections.CleanUp()

	connections.NewMySqlConnection()
	connections.NewRedisConnection()

	server.Start()
}
