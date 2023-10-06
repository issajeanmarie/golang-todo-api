package todos

import (
	"todo-api/api/handler"

	"github.com/gin-gonic/gin"
)


func Post(r *gin.Engine) {
	r.POST("/todos", handler.PostTodosHandler)
}