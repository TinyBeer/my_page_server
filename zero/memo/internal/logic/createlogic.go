package logic

import (
	"context"
	"database/sql"

	"personal_page/zero/memo/internal/svc"
	"personal_page/zero/memo/internal/types"
	"personal_page/zero/model/memo"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.MemoCreateReq) error {
	_, err := l.svcCtx.MemoModel.Insert(l.ctx, &memo.GormMemo{
		Id:      0,
		Content: req.Content,
		DeletedAt: sql.NullInt64{
			Int64: 0,
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
