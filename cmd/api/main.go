package main

import (
	"github.com/afiffaizun/todo-app-cicd/internal/handler"
	"github.com/afiffaizun/todo-app-cicd/internal/model"
	"github.com/afiffaizun/todo-app-cicd/internal/repository"
	"github.com/afiffaizun/todo-app-cicd/internal/service"
	"github.com/afiffaizun/todo-app-cicd/pkg/config"
	"github.com/afiffaizun/todo-app-cicd/pkg/database"
	"github.com/afiffaizun/todo-app-cicd/pkg/utils"
	"github.com/gin-gonic/gin"
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

	db.AutoMigrate(&model.Todo{}, &model.User{})

	// Initialize JWT utility
	jwtUtil := utils.NewJWTUtil(cfg.JWTSecret)

	// Initialize repositories
	todoRepo := repository.NewTodoRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	todoHandler := handler.NewTodoHandler(todoRepo)
	authService := service.NewAuthService(userRepo, jwtUtil)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/profile", authHandler.GetProfile)
		}

		// Todo routes
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos", todoHandler.GetAllTodos)
		api.PUT("/todos/:id", todoHandler.UpdateTodo)
		api.DELETE("/todos/:id", todoHandler.DeleteTodo)
	}

	r.Run(":" + cfg.ServerPort)
}
