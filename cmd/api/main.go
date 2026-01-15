package main

import (
	"github.com/gin-gonic/gin"
	"github.com/afiffaizun/todo-app-cicd/internal/model"
	"github.com/afiffaizun/todo-app-cicd/internal/handler"
	"github.com/afiffaizun/todo-app-cicd/internal/repository"
	"github.com/afiffaizun/todo-app-cicd/pkg/database"
    "github.com/afiffaizun/todo-app-cicd/pkg/config"

)

func main() {
    cfg := config.LoadConfig()
    
    db := database.NewPostgresDB(
        cfg.DBHost,
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBName,
        cfg.DBPort,
    )
    
    db.AutoMigrate(&model.Todo{})
    
    todoRepo := repository.NewTodoRepository(db)
    todoHandler := handler.NewTodoHandler(todoRepo)
    
    r := gin.Default()
    
    api := r.Group("/api/v1")
    {
        api.POST("/todos", todoHandler.CreateTodo)
        api.GET("/todos", todoHandler.GetAllTodos)
        api.PUT("/todos/:id", todoHandler.UpdateTodo)
        api.DELETE("/todos/:id", todoHandler.DeleteTodo)
    }
    
    r.Run(":" + cfg.ServerPort)
}