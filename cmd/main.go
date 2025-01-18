package main

import (
	"mock-crud-api/config"
	"mock-crud-api/routes"
)

func main() {
	config.ConnectDatabase() // Initialize DB
	router := routes.SetupRouter()
	router.Run(":8080")
}
