package router

import (
	"fmt"
	todoHandler "todo-api/api/handler"
	consts "todo-api/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	fmt.Println("Running on: ", consts.URL)
	r.GET(consts.ToDosURL, todoHandler.GetTodos)
	r.POST(consts.ToDosURL, todoHandler.AddTodo)

	return r
}