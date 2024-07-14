package handler

import (
	"net/http"

	"personal_page/domain"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	uc domain.TodoUsecase
}

func NewTodoHandler(uc domain.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		uc: uc,
	}
}

// List 获取TODO列表接口
// @Summary 获取TODO列表资源
// @Description 使用access_token获取TODO列表资源
// @Tags TODO相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} domain.TodoListResp
// @Router /todo [get]
func (h *TodoHandler) List(ctx *gin.Context) {
	list, err := h.uc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
	}

	data := make([]*domain.Todo, 0, len(list))
	for _, v := range list {
		todo := &domain.Todo{
			ID:               v.ID,
			Content:          v.Content,
			Times:            v.Times,
			CompletetedTimes: v.CompletedTimes,
			Duration:         v.Duration,
		}
		if v.Completed {
			todo.CompletedAt = v.CompletedAt.Format("2006-01-02 15:03:04")
		}
		data = append(data, todo)
	}
	ctx.JSON(http.StatusOK, domain.TodoListResp{
		BaseResp: domain.BaseResp{Status: domain.StatusOk},
		Data:     data,
	})
}

// Create 创建TODO
// @Summary 根据内容创建新的TODO
// @Description 使用access_token创建TODO
// @Tags TODO相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body domain.TodoCreateReq true "TODO创建参数"
// @Success 200 {object} domain.BaseResp
// @Router /todo [post]
func (h *TodoHandler) Create(ctx *gin.Context) {
	var req domain.TodoCreateReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}

	todo := &domain.UcTodo{
		Content:  req.Content,
		Duration: req.Duration,
		Times:    req.Times,
	}
	if todo.Times == 0 {
		todo.Times = 1
	}
	err = h.uc.Create(ctx, todo)
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

// DeleteByID 删除TOOD
// @Summary 根据Id删除TOOD
// @Description 使用access_token删除指定id的TODO
// @Tags TODO相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body domain.TodoDeleteReq true "TODO删除参数"
// @Success 200 {object} domain.BaseResp
// @Failed 200 {object} domain.BaseResp
// @Router /todo [delete]
func (h *TodoHandler) DeleteByID(ctx *gin.Context) {
	var req domain.TodoDeleteReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = h.uc.DeleteByID(ctx, req.ID)
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

// Complete 完成TODO
// @Summary 根据ID更新TODO完成时间
// @Description 使用access_token根据ID更新TODO完成时间
// @Tags TODO相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body domain.TodoCompleteReq true "TODO完成参数"
// @Success 200 {object} domain.BaseResp
// @Failed 200 {object} domain.BaseResp
// @Router /todo [put]
func (h *TodoHandler) CompleteByID(ctx *gin.Context) {
	var req domain.TodoCompleteReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.uc.CompleteByID(ctx, req.ID)
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
