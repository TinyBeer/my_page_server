package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type TodoListResp struct {
	BaseResp
	Data []*Todo
}

type TodoCreateReq struct {
	Content string `json:"content,omitempty"`
}

type TodoDeleteReq struct {
	ID string `json:"id,omitempty"`
}

type UcTodo struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
}

type TodoUsecase interface {
	List(ctx context.Context) ([]*UcTodo, error)
	Create(ctx context.Context, todo *UcTodo) error
	DeleteByID(ctx context.Context, id uint) error
}

type RepoTodo struct {
	gorm.Model
	Content     string
	CompletedAt int64
}

func (*RepoTodo) TableName() string {
	return "todo"
}

type TodoRepository interface {
	List(ctx context.Context) ([]*RepoTodo, error)
	Create(ctx context.Context, todo *RepoTodo) error
	DeleteByID(ctx context.Context, id uint) error
}
