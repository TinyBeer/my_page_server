package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID               string `json:"id,omitempty"`
	Content          string `json:"content,omitempty"`
	CompletedAt      string `json:"completed_at,omitempty"`
	Times            int    `json:"times,omitempty"`
	CompletetedTimes int    `json:"completeted_times,omitempty"`
	Duration         int64  `json:"duration,omitempty"`
}

type TodoListResp struct {
	BaseResp
	Data []*Todo
}

type TodoCreateReq struct {
	Content  string `json:"content,omitempty"`
	Times    int    `json:"times,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}

type TodoDeleteReq struct {
	ID string `json:"id,omitempty"`
}

type TodoCompleteReq struct {
	ID string `json:"id,omitempty"`
}

type UcTodo struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Completed      bool
	CompletedAt    time.Time
	Content        string
	Times          int
	CompletedTimes int
	Duration       int64
}

type TodoUsecase interface {
	List(ctx context.Context) ([]*UcTodo, error)
	Create(ctx context.Context, todo *UcTodo) error
	DeleteByID(ctx context.Context, id string) error
	CompleteByID(ctx context.Context, id string) error
}

type RepoTodo struct {
	gorm.Model
	Content        string
	CompletedAt    int64
	Times          int
	CompletedTimes int
	Duration       int64
}

func (*RepoTodo) TableName() string {
	return "todo"
}

type TodoRepository interface {
	List(ctx context.Context) ([]*RepoTodo, error)
	GetByID(ctx context.Context, id uint) (*RepoTodo, error)
	Create(ctx context.Context, todo *RepoTodo) error
	DeleteByID(ctx context.Context, id uint) error
	UpdateByID(ctx context.Context, todo *RepoTodo) error
}
