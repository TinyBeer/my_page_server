package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          string `json:"id,omitempty"`
	Content     string `json:"content,omitempty"`
	CompletedAt string `json:"completed_at,omitempty"`
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

type TodoCompleteReq struct {
	ID string `json:"id,omitempty"`
}

type UcTodo struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Completed   bool
	CompletedAt time.Time
	Content     string
}

type TodoUsecase interface {
	List(ctx context.Context) ([]*UcTodo, error)
	Create(ctx context.Context, todo *UcTodo) error
	DeleteByID(ctx context.Context, id string) error
	CompleteByID(ctx context.Context, id string) error
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
	CompleteByID(ctx context.Context, id uint) error
}
