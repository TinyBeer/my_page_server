package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"personal_page/zero/memo/internal/logic"
	"personal_page/zero/memo/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}