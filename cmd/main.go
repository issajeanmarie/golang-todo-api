package main

import (
	"todo-api/api/router"
)

func main() {
	r := router.SetupRouter()
	r.Run();
}