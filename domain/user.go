package domain

import (
	"context"

	"gorm.io/plugin/soft_delete"
)

type RepoUser struct {
	ID        int                   `gorm:"size:10"`
	Name      string                `gorm:"size:10;not null"`
	Password  string                `gorm:"size:60;not null"`
	CreatedAt int64                 `gorm:"autoCreateTime"`
	UpdatedAt int64                 `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
}

func (*RepoUser) TableName() string {
	return "user"
}

type UserRepository interface {
	Create(ctx context.Context, user *RepoUser) error
	Check(ctx context.Context, user *RepoUser) error
	GetByName(ctx context.Context, name string) (*RepoUser, error)
}

type UcUser struct {
	ID       int
	Name     string
	Password string
}
