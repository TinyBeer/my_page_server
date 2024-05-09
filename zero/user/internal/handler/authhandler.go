package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"personal_page/zero/user/internal/logic"
	"personal_page/zero/user/internal/svc"
)

func AuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthLogic(r.Context(), svcCtx)
		err := l.Auth()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
