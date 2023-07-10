package main

import (
	"crud-api/database"
	"crud-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
