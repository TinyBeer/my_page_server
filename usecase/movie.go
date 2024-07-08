package usecase

import (
	"context"

	"personal_page/model"
	"personal_page/repository"
)

type MovieUc struct {
	repo repository.MovieRepository
}

// Create implements MovieUsecase.
func (m *MovieUc) Create(ctx context.Context, movie *model.Movie) error {
	return nil
}

// DeleteById implements MovieUsecase.
func (m *MovieUc) DeleteById(ctx context.Context, id string) error {
	return nil
}

// List implements MovieUsecase.
func (m *MovieUc) List(ctx context.Context) ([]*model.Movie, error) {
	return nil, nil
}

func NewMovieUc(repo repository.MovieRepository) *MovieUc {
	return &MovieUc{
		repo: repo,
	}
}

var _ MovieUsecase = new(MovieUc)
