package router

import (
	"os"
	todoHandler "todo-api/api/handler"
	consts "todo-api/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	


	port := os.Getenv("PORT"); 
	if port == ""{
		port = consts.URL
	}
	
	r.GET(consts.ToDosURL, todoHandler.GetTodos)
	r.POST(consts.ToDosURL, todoHandler.AddTodo)
	r.DELETE(consts.ToDosURL + "/:id", todoHandler.DeleteTodo)

	return r
}