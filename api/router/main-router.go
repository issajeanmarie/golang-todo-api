package router

import (
	"todo-api/api/router/todos"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	todos.Get(r)
	todos.Post(r)

	return r
}