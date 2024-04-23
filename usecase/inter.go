package usecase

import (
	"context"
	"personal_page/model"
)

type UserUsecase interface {
	Login(ctx context.Context, username string, password string) (*model.LoginResult, error)
	CheckAccessToken(ctx context.Context, accessToken string) (*model.User, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
}
