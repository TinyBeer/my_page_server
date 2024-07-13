package mysql

import (
	"context"

	"personal_page/domain"

	"gorm.io/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

// Create implements domain.TodoRepository.
func (t *TodoRepo) Create(ctx context.Context, todo *domain.RepoTodo) error {
	return t.db.WithContext(ctx).Create(todo).Error
}

// DeleteByID implements domain.TodoRepository.
func (t *TodoRepo) DeleteByID(ctx context.Context, id uint) error {
	return t.db.WithContext(ctx).Delete(&domain.RepoTodo{Model: gorm.Model{ID: id}}).Error
}

// List implements domain.TodoRepository.
func (t *TodoRepo) List(ctx context.Context) ([]*domain.RepoTodo, error) {
	var list []*domain.RepoTodo
	err := t.db.WithContext(ctx).Find(&list).Error
	return list, err
}

func NewTodoRepo(db *gorm.DB) domain.TodoRepository {
	repo := &TodoRepo{
		db: db,
	}
	err := repo.db.AutoMigrate(&domain.RepoTodo{})
	if err != nil {
		panic(err)
	}
	return repo
}
