package main

import (
	"fmt"
	"todo-api/api/router"
	"todo-api/config"
)

func main() {

	r := router.Router()
	port := config.PORT()

	r.Run(port)
	fmt.Println("Running on port: 3000")
}