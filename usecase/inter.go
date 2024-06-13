package usecase

import (
	"context"

	"personal_page/model"
)

type UserUsecase interface {
	Login(ctx context.Context, username, password string) (*model.LoginResult, error)
	CheckAccessToken(ctx context.Context, accessToken string) (*model.User, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
}

type MemoUsecase interface {
	List(ctx context.Context) ([]*model.Memo, error)
	Create(ctx context.Context, content string) error
	CompleteWithId(ctx context.Context, id string) error
	DeleteById(ctx context.Context, id string) error
}

type MovieUsecase interface {
	List(ctx context.Context) ([]*model.Movie, error)
	Create(ctx context.Context, movie *model.Movie) error
	DeleteById(ctx context.Context, id string) error
}
