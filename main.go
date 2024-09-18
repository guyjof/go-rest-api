package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Build a RESTful API in Go", Completed: false},
	{ID: "3", Item: "Build a React app", Completed: false},
}

func getTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, newTodo)

	context.JSON(http.StatusCreated, newTodo)
}

func getTodoByID(id string) (*Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	context.JSON(http.StatusOK, todo)
}

func updateTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var updatedTodo Todo
	if err := context.BindJSON(&updatedTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Item = updatedTodo.Item
	todo.Completed = updatedTodo.Completed

	context.JSON(http.StatusOK, todo)
}

func deleteTodoByID(id string) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	err := deleteTodoByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run(":8080")
}
