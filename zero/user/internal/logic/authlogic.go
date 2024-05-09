package logic

import (
	"context"
	"fmt"

	"personal_page/zero/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth() error {
	fmt.Println(l.ctx.Value("username"))
	return nil
}
