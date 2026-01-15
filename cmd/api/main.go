package main

import (
	"github.com/gin-gonic/gin"
	"github.com/afiffaizun/todo-app-cicd/internal/model"
	"github.com/afiffaizun/todo-app-cicd/internal/handler"
	"github.com/afiffaizun/todo-app-cicd/internal/repository"
	"github.com/afiffaizun/todo-app-cicd/pkg/database"
)

func main() {
	db := database.NewPostgresDB("localhost", "todouser", "todopass", "tododb", "5432")

	// Auto migrate
    db.AutoMigrate(&model.Todo{})
    
    // Setup handler
    todoRepo := repository.NewTodoRepository(db)
    todoHandler := handler.NewTodoHandler(todoRepo)
    
    r := gin.Default()
    
    // Routes
    api := r.Group("/api/v1")
    {
        api.POST("/todos", todoHandler.CreateTodo)
        api.GET("/todos", todoHandler.GetAllTodos)
        api.PUT("/todos/:id", todoHandler.UpdateTodo)
        api.DELETE("/todos/:id", todoHandler.DeleteTodo)
    }

	r.Run(":8080")
}