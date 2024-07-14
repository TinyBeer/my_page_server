package usecase

import (
	"context"
	"errors"

	"personal_page/domain"
	"personal_page/usecase/auth"

	"golang.org/x/crypto/bcrypt"
)

type TokenUsecase struct {
	userRepo domain.UserRepository
}

// AccessTokenCheck implements domain.TokenUsecase.
func (t *TokenUsecase) AccessTokenCheck(ctx context.Context, accessToken string) (*domain.UcUser, error) {
	claims, err := auth.ParseAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	u, err := t.userRepo.GetByName(ctx, claims.UserName)
	if err != nil {
		return nil, err
	}
	return &domain.UcUser{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
	}, nil
}

// AccessTokenGen implements domain.TokenUsecase.
func (t *TokenUsecase) AccessTokenGen(ctx context.Context, refreshToken string) (string, error) {
	claims, err := auth.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	return auth.GenerateAccessToken(claims.UserName)
}

// TokenGen implements domain.TokenUsecase.
func (t *TokenUsecase) TokensGen(ctx context.Context, user *domain.UcUser) (string, string, error) {
	u, err := t.userRepo.GetByName(ctx, user.Name)
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return "", "", errors.New("账号或密码错误")
	}

	accessToken, err := auth.GenerateAccessToken(user.Name)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := auth.GenerateRefreshToken(user.Name)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func NewTokenUsecase(repo domain.UserRepository) domain.TokenUsecase {
	return &TokenUsecase{
		userRepo: repo,
	}
}
