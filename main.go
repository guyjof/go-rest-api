package main

import (
	"go-rest-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", handlers.GetTodos)
	router.POST("/todos", handlers.AddTodo)
	router.GET("/todos/:id", handlers.GetTodo)
	router.PATCH("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)
	router.Run(":8080")
}
