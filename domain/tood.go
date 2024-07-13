package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	StatusError = "error"
	StatusOk    = "ok"
)

type DeliTodo struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type BaseResp struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type DeliTodoListResp struct {
	BaseResp
	Data []*DeliTodo
}

type DeliTodoCreateReq struct {
	Content string `json:"content,omitempty"`
}

type DeliTodoDeleteReq struct {
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
