package main

import (
	"simple-crud-task/connection"
	"simple-crud-task/handlers"
)

func main() {
	connection.ConnectDB()

	handlers.Handlers()
}
