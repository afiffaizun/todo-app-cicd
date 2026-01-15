package repository

import (
	"github.com/afiffaizun/todo-app-cicd/internal/model"
    "gorm.io/gorm"
)

type TodoRepository interface {
    Create(todo *model.Todo) error
    FindAll() ([]model.Todo, error)
    FindByID(id uint) (*model.Todo, error)
    Update(todo *model.Todo) error
    Delete(id uint) error
}

type todoRepository struct {
    db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
    return &todoRepository{db}
}

func (r *todoRepository) Create(todo *model.Todo) error {
    return r.db.Create(todo).Error
}

func (r *todoRepository) FindAll() ([]model.Todo, error) {
    var todos []model.Todo
    err := r.db.Find(&todos).Error
    return todos, err
}

func (r *todoRepository) FindByID(id uint) (*model.Todo, error) {
    var todo model.Todo
    err := r.db.First(&todo, id).Error
    return &todo, err
}

func (r *todoRepository) Update(todo *model.Todo) error {
    return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
    return r.db.Delete(&model.Todo{}, id).Error
}