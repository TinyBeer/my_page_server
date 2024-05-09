package logic

import (
	"context"

	"personal_page/zero/memo/internal/svc"
	"personal_page/zero/memo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() (resp *types.MemoListResp, err error) {
	memoes, err := l.svcCtx.MemoModel.Find(l.ctx)
	if err != nil {
		return nil, err
	}

	resp = &types.MemoListResp{
		Memoes: nil,
	}
	for _, item := range memoes {
		resp.Memoes = append(resp.Memoes, types.MemoItem{
			Id:      uint(item.Id),
			Content: item.Content,
		})
	}

	return
}
