package middleware

import (
	"net/http"
	"strings"

	"personal_page/domain"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	token domain.TokenUsecase
}

func NewMiddleware(token domain.TokenUsecase) *Middleware {
	return &Middleware{
		token: token,
	}
}

func (m *Middleware) JWT(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")
	if accessToken == "" {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: domain.ErrorNoToken,
		})
		ctx.Abort()
	}

	user, err := m.token.AccessTokenCheck(ctx, accessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Set("user", user)
}
