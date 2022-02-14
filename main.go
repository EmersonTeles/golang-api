package main

import (
	"fmt"

	"github.com/EmersonTeles/golang-api/database"
	"github.com/EmersonTeles/golang-api/routes"
)

func main() {
	database.ConnectToDatabase()
	fmt.Println("iniciando server on 8000")
	routes.HandleRequest()
}
