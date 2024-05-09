package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"personal_page/zero/memo/internal/logic"
	"personal_page/zero/memo/internal/svc"
	"personal_page/zero/memo/internal/types"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemoDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteLogic(r.Context(), svcCtx)
		err := l.Delete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
