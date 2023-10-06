package handler

import (
	"net/http"
	"todo-api/api/module"
	"todo-api/config"

	"github.com/gin-gonic/gin"
)

func PostTodosHandler(ctx *gin.Context) {
	var newTodo module.Todo
	if err := ctx.BindJSON(&newTodo); err != nil{
		return;
	}

	config.CUSTOM_TODOS()
	ctx.IndentedJSON(http.StatusCreated, newTodo)

}