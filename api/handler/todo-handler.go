package todoHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID int `json:"id"`
	Label string `json:"label"`
	Completed bool `json:"completed"`
}

var todoList = []Todo{
	{
		ID: 1,
		Label: "Cook",
		Completed: false,
	},
}

func GetTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todoList)
}

func AddTodo(context *gin.Context){
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil{
		return
	}
	
	todoList = append(todoList, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func DeleteTodo(context *gin.Context){
	todoID := context.Param("id");

	id, _ := strconv.Atoi(todoID)
	if todoID == "" {
		context.JSON(http.StatusBadRequest, "Invalid ID");
		return;
	}

	indexToDelete := -1
	for i, todo := range todoList{
		if todo.ID == id{
			indexToDelete = i
		}
	}

	if(indexToDelete == -1){
		context.JSON(http.StatusBadRequest, "Todo not found!")
		return
	}

	todoList = append(todoList[:indexToDelete], todoList[indexToDelete+1:]...)
	context.JSON(http.StatusOK, gin.H{"message":"Todo deleted successfully!"})
}