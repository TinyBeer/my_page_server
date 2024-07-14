package handler

import (
	"context"
	"net/http"
	"strings"

	"personal_page/domain"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	token domain.TokenUsecase
}

// Gen 用户登录接口
// @Summary 用于用户登录
// @Description 完成用户登陆操作分发access_token和refresh_token
// @Tags TOKEN相关接口
// @Accept application/json
// @Produce application/json
// @Param {object} body domain.TokenReq true "登录参数"
// @Success 200 {object} domain.TokenResp
// @Router /token [post]
func (h *TokenHandler) Gen(ctx *gin.Context) {
	var req domain.TokenReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.TokenResp{
			BaseResp: domain.BaseResp{
				Status:  domain.StatusError,
				Message: err.Error(),
			},
		})
		return
	}

	at, rt, err := h.token.TokensGen(context.Background(), &domain.UcUser{
		Name:     req.Username,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, domain.TokenResp{
			BaseResp: domain.BaseResp{
				Status:  domain.StatusError,
				Message: err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.TokenResp{
		BaseResp: domain.BaseResp{
			Status: domain.StatusOk,
		},
		AccessToken:  at,
		RefreshToken: rt,
	})
}

// Check 验证access_token
// @Summary 验证用户身份
// @Description 使用access_token验证用户身份
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} domain.BaseResp
// @Router /token/check [get]
func (h *TokenHandler) Check(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, domain.BaseResp{
		Status: domain.StatusOk,
	})
}

// AccessToken 重新获取访问令牌
// @Summary 更新访问令牌
// @Description 使用refresh_token更新访问令牌
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "refresh_token"
// @Success 200 {object} domain.AccessTokenResp
// @Router /token/access_token [post]
func (h *TokenHandler) AccessToken(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("Authorization")
	refreshToken = strings.TrimPrefix(refreshToken, "Bearer ")
	if refreshToken == "" {
		ctx.JSON(http.StatusUnauthorized, "no token")
		return
	}

	accessToken, err := h.token.AccessTokenGen(context.Background(), refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, domain.AccessTokenResp{
		BaseResp: domain.BaseResp{
			Status: domain.StatusOk,
		},
		AccessToken: accessToken,
	})
}

func NewTokenHandler(token domain.TokenUsecase) *TokenHandler {
	return &TokenHandler{
		token: token,
	}
}
