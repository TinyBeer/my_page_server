package usecase

import (
	"context"
	"errors"

	"personal_page/model"
	"personal_page/repository"
	"personal_page/usecase/auth"

	"golang.org/x/crypto/bcrypt"
)

type UserUc struct {
	userRepo repository.UserRepository
}

// CheckAccessToken implements UserUsecase.
func (uc *UserUc) CheckAccessToken(ctx context.Context, accessToken string) (*model.User, error) {
	claims, err := auth.ParseAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	return uc.userRepo.GetUserByName(claims.UserName)
}

// Login implements UserUsecase.
func (uc *UserUc) Login(ctx context.Context, username, password string) (*model.LoginResult, error) {
	u, err := uc.userRepo.GetUserByName(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, errors.New("账号或密码错误")
	}

	accessToken, err := auth.GenerateAccessToken(username)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(username)
	if err != nil {
		return nil, err
	}
	return &model.LoginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshAccessToken implements UserUsecase.
func (uc *UserUc) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := auth.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	return auth.GenerateAccessToken(claims.UserName)
}

func NewUserUc(userRepo repository.UserRepository) *UserUc {
	return &UserUc{
		userRepo: userRepo,
	}
}

var _ UserUsecase = &UserUc{}
