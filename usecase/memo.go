package usecase

import (
	"context"

	"personal_page/model"
	"personal_page/repository"
)

type MemoUc struct {
	repo repository.MemoRepository
}

func NewMemoUc(repo repository.MemoRepository) *MemoUc {
	return &MemoUc{
		repo: repo,
	}
}

// CompleteWithId implements MemoUsecase.
func (m *MemoUc) CompleteWithId(ctx context.Context, id uint) error {
	return m.repo.CompleteWithId(id)
}

// Create implements MemoUsecase.
func (m *MemoUc) Create(ctx context.Context, content string) error {
	return m.repo.CreateMemo(content)
}

// DeleteById implements MemoUsecase.
func (m *MemoUc) DeleteById(ctx context.Context, id uint) error {
	return m.repo.DeleteMemoById(id)
}

// List implements MemoUsecase.
func (m *MemoUc) List(ctx context.Context) ([]*model.Memo, error) {
	return m.repo.ListMemo()
}

var _ MemoUsecase = new(MemoUc)
