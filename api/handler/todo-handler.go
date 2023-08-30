package todoHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Label string `json:"label"`
	Completed bool `json:"completed"`
}

var todoList = []Todo{
	{
		Label: "Cook",
		Completed: false,
	},
}

func GetTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todoList)
}

func AddTodo(context*gin.Context){
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil{
		return
	}
	
	todoList = append(todoList, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}