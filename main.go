package main

import (
	"github.com/Ismadrade/api-go-gin/database"
	"github.com/Ismadrade/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
