package todos

import (
	"todo-api/api/handler"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine) {
	r.GET("/todos", handler.GetTodosHandler)
}