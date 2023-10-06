package config

import (
	"todo-api/api/module"
)

func PORT() string {
	return "localhost:3000"
}

func CUSTOM_TODOS() []module.Todo{
	todos := []module.Todo {
		{
			Name: "Cook",
		},{
			Name: "Call mom",
		},
		
	}
	return todos
}