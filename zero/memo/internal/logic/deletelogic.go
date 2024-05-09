package logic

import (
	"context"

	"personal_page/zero/memo/internal/svc"
	"personal_page/zero/memo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.MemoDeleteReq) error {
	return l.svcCtx.MemoModel.Delete(l.ctx, int64(req.Id))
}
