package usecase

import (
	"context"

	"personal_page/domain"
)

type TodoUsecase struct {
	repo domain.TodoRepository
}

// Create implements domain.TodoUsecase.
func (t *TodoUsecase) Create(ctx context.Context, todo *domain.UcTodo) error {
	return t.repo.Create(ctx, &domain.RepoTodo{
		Content: todo.Content,
	})
}

// DeleteByID implements domain.TodoUsecase.
func (t *TodoUsecase) DeleteByID(ctx context.Context, id uint) error {
	return t.repo.DeleteByID(ctx, id)
}

// List implements domain.TodoUsecase.
func (t *TodoUsecase) List(ctx context.Context) ([]*domain.UcTodo, error) {
	list, err := t.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*domain.UcTodo, 0, len(list))
	for _, v := range list {
		result = append(result, &domain.UcTodo{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Content:   v.Content,
		})
	}
	return result, nil
}

func NewTodoUsecase(repo domain.TodoRepository) domain.TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}
