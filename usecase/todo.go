package usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	"personal_page/domain"
)

type TodoUsecase struct {
	repo domain.TodoRepository
}

// CompleteByID implements domain.TodoUsecase.
func (t *TodoUsecase) CompleteByID(ctx context.Context, id string) error {
	iid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	todo, err := t.repo.GetByID(ctx, uint(iid))
	if err != nil {
		return err
	}

	if todo.CompletedAt != 0 {
		return errors.New("todo[" + id + "] already completed")
	}
	todo.CompletedTimes++
	if todo.CompletedTimes == todo.Times {
		todo.CompletedAt = time.Now().Unix()
	}

	return t.repo.UpdateByID(ctx, todo)
}

// Create implements domain.TodoUsecase.
func (t *TodoUsecase) Create(ctx context.Context, todo *domain.UcTodo) error {
	return t.repo.Create(ctx, &domain.RepoTodo{
		Content:  todo.Content,
		Times:    todo.Times,
		Duration: todo.Duration,
	})
}

// DeleteByID implements domain.TodoUsecase.
func (t *TodoUsecase) DeleteByID(ctx context.Context, id string) error {
	iid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return t.repo.DeleteByID(ctx, uint(iid))
}

// List implements domain.TodoUsecase.
func (t *TodoUsecase) List(ctx context.Context) ([]*domain.UcTodo, error) {
	list, err := t.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*domain.UcTodo, 0, len(list))
	for _, v := range list {
		todo := &domain.UcTodo{
			ID:             strconv.Itoa(int(v.ID)),
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			Completed:      v.CompletedAt != 0,
			Content:        v.Content,
			Times:          v.Times,
			CompletedTimes: v.CompletedTimes,
			Duration:       v.Duration,
		}
		if todo.Completed {
			todo.CompletedAt = time.Unix(v.CompletedAt, 0)
		}
		result = append(result, todo)
	}
	return result, nil
}

func NewTodoUsecase(repo domain.TodoRepository) domain.TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}
