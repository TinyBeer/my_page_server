package middleware

import (
	"net/http"

	"personal_page/usecase"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	uc usecase.UserUsecase
}

func NewMiddleware(uc usecase.UserUsecase) *Middleware {
	return &Middleware{
		uc: uc,
	}
}

func (m *Middleware) JWT(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")
	if accessToken == "" {
		ctx.JSON(http.StatusOK, "no token")
		ctx.Abort()
	}

	user, err := m.uc.CheckAccessToken(ctx, accessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
}
