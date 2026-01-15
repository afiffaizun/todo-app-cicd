package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/afiffaizun/todo-app-cicd/internal/model"
	"github.com/afiffaizun/todo-app-cicd/internal/repository"
)

type TodoHandler struct {
    repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
    return &TodoHandler{repo}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
    var todo model.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := h.repo.Create(&todo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
    todos, err := h.repo.FindAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    todo, err := h.repo.FindByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    
    if err := c.ShouldBindJSON(todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := h.repo.Update(todo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    if err := h.repo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}