package handlers

import (
	"errors"
	"net/http"

	"go-rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTodos(context *gin.Context) {
	context.JSON(http.StatusOK, models.Todos)
}

func AddTodo(context *gin.Context) {
	var newTodo models.Todo
	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique ID for the new todo
	newTodo.ID = uuid.New().String()

	models.Todos = append(models.Todos, newTodo)
	context.JSON(http.StatusCreated, newTodo)
}

func GetTodoByID(id string) (*models.Todo, error) {
	for i, todo := range models.Todos {
		if todo.ID == id {
			return &models.Todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	context.JSON(http.StatusOK, todo)
}

func UpdateTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoByID(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var updatedTodo models.Todo
	if err := context.BindJSON(&updatedTodo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Item = updatedTodo.Item
	todo.Completed = updatedTodo.Completed

	context.JSON(http.StatusOK, todo)
}

func DeleteTodoByID(id string) error {
	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos = append(models.Todos[:i], models.Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	err := DeleteTodoByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
