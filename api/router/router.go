package router

import (
	"os"
	todoHandler "todo-api/api/handler"
	consts "todo-api/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()


	port := os.Getenv("PORT"); 
	if port == ""{
		port = consts.URL
	}
	
	r.GET(consts.ToDosURL, todoHandler.GetTodos)
	r.POST(consts.ToDosURL, todoHandler.AddTodo)

	return r
}