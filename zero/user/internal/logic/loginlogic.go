package logic

import (
	"context"
	"database/sql"

	"personal_page/zero/user/internal/common"
	"personal_page/zero/user/internal/svc"
	"personal_page/zero/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByNameDeletedAt(l.ctx, req.Username, sql.NullInt64{
		Int64: 0,
		Valid: true,
	})
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	acToken, err := common.GenerateAccessToken(user.Name)
	if err != nil {
		return nil, err
	}
	rfToken, err := common.GenerateRefreshToken(user.Name)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  acToken,
		RefreshToken: rfToken,
	}, nil
}
