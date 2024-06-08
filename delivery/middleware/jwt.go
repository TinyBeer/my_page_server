package middleware

import (
	"net/http"
	"strings"

	"personal_page/model"
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
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")
	if accessToken == "" {
		ctx.JSON(http.StatusOK, model.Base{
			Status:  model.StatusError,
			Message: model.ErrorNoToken,
		})
		ctx.Abort()
	}

	user, err := m.uc.CheckAccessToken(ctx, accessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.Base{
			Status:  model.StatusError,
			Message: err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
}
