package handler

import (
	"net/http"
	"todo-api/config"

	"github.com/gin-gonic/gin"
)

func GetTodosHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, config.CUSTOM_TODOS())
}