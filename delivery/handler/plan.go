package handler

import (
	"net/http"

	"personal_page/domain"

	"github.com/gin-gonic/gin"
)

type PlanHandler struct {
	plan domain.PlanUsecase
}

func NewPlanHandler(plan domain.PlanUsecase) *PlanHandler {
	return &PlanHandler{
		plan: plan,
	}
}

// List 获取Plan列表接口
// @Summary 获取Plan列表资源
// @Description 使用access_token获取Plan列表资源
// @Tags Plan相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} domain.PlanListResp
// @Router /plan [get]
func (h *PlanHandler) List(ctx *gin.Context) {
	list, err := h.plan.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
	}

	data := make([]*domain.Plan, 0, len(list))
	for _, v := range list {
		plan := &domain.Plan{
			ID:       v.ID,
			Type:     v.Type,
			Content:  v.Content,
			Times:    v.Times,
			Duration: v.Duration,
		}

		data = append(data, plan)
	}
	ctx.JSON(http.StatusOK, domain.PlanListResp{
		BaseResp: domain.BaseResp{Status: domain.StatusOk},
		Data:     data,
	})
}

// Create 创建Plan
// @Summary 根据内容创建新的Plan
// @Description 使用access_token创建Plan
// @Tags Plan相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body domain.PlanCreateReq true "TODO创建参数"
// @Success 200 {object} domain.BaseResp
// @Router /plan [post]
func (h *PlanHandler) Create(ctx *gin.Context) {
	var req domain.PlanCreateReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}

	plan := &domain.UcPlan{
		Content:  req.Content,
		Duration: req.Duration,
		Times:    req.Times,
	}
	if plan.Times == 0 {
		plan.Times = 1
	}
	err = h.plan.Create(ctx, plan)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.BaseResp{
		Status: domain.StatusOk,
	})
}

// DeleteByID 删除Plan
// @Summary 根据Id删除Plan
// @Description 使用access_token删除指定id的Plan
// @Tags TODO相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body domain.PlanDeleteReq true "TODO删除参数"
// @Success 200 {object} domain.BaseResp
// @Failed 200 {object} domain.BaseResp
// @Router /plan [delete]
func (h *PlanHandler) DeleteByID(ctx *gin.Context) {
	var req domain.PlanDeleteReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = h.plan.DeleteByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.BaseResp{
		Status: domain.StatusOk,
	})
}
