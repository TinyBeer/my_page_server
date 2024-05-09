package logic

import (
	"context"

	"personal_page/zero/user/internal/common"
	"personal_page/zero/user/internal/svc"
	"personal_page/zero/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshReq) (resp *types.RefreshResp, err error) {
	claims, err := common.ParseRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	acToken, err := common.GenerateAccessToken(claims.Username)
	if err != nil {
		return nil, err
	}

	return &types.RefreshResp{
		AccessToken: acToken,
	}, nil
}
