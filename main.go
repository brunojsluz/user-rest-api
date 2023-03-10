package main

import (
	"user-rest-api/database"
	"user-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
