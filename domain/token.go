package domain

import (
	"context"
)

type TokenReq struct {
	Username string `json:"username,omitempty" binding:"min=4,max=10"`
	Password string `json:"password,omitempty" binding:"min=3,max=18"`
}

type TokenResp struct {
	BaseResp
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AccessTokenResp struct {
	BaseResp
	AccessToken string `json:"access_token,omitempty"`
}

type TokenUsecase interface {
	TokensGen(ctx context.Context, user *UcUser) (accessToken string, refreshToken string, err error)
	AccessTokenCheck(ctx context.Context, accessToken string) (*UcUser, error)
	AccessTokenGen(ctx context.Context, refreshToken string) (accessToken string, err error)
}
