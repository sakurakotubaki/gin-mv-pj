package controller

import (
	"go-mc-app/model"

	"gorm.io/gorm"
)

// TodoControllerは、Todoに関するリクエストを処理するための構造体です。
type TodoController struct {
    DB *gorm.DB
}

// NewTodoControllerは、TodoControllerを初期化して返します。
func NewTodoController(db *gorm.DB) *TodoController {
    return &TodoController{
        DB: db,
    }
}

// CreateTodoは、Todoを作成します。
func (t *TodoController) CreateTodo(todo *model.Todo) error {
    if err := t.DB.Create(todo).Error; err != nil {
        return err
    }
    return nil
}

// GetTodosは、Todoを取得します。
func (t *TodoController) GetTodos() ([]model.Todo, error) {
    var todos []model.Todo
    if err := t.DB.Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}

// UpdateTodoは、Todoを更新します。
func (t *TodoController) UpdateTodo(todo *model.Todo) error {
    if err := t.DB.Save(todo).Error; err != nil {
        return err
    }
    return nil
}

// DeleteTodoは、Todoを削除します。
func (t *TodoController) DeleteTodo(todo *model.Todo) error {
    if err := t.DB.Delete(todo).Error; err != nil {
        return err
    }
    return nil
}