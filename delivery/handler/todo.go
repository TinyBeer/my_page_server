package handler

import (
	"net/http"
	"strconv"

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
		data = append(data, &domain.Todo{
			ID:      strconv.Itoa(int(v.ID)),
			Content: v.Content,
		})
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
	err = h.uc.Create(ctx, &domain.UcTodo{Content: req.Content})
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.BaseResp{
		Status:  domain.StatusOk,
		Message: "",
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
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, domain.BaseResp{
			Status:  domain.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = h.uc.DeleteByID(ctx, uint(id))
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

// // CompleteWithId 备忘录完成
// // @Summary 根据Id更新备忘录完成状态
// // @Description 使用access_token获取视频列表资源
// // @Tags 备忘录相关接口
// // @Accept application/json
// // @Produce application/json
// // @Param Authorization header string true "访问令牌"
// // @Param {object} body model.MemoCompleteRequest true "备忘录完成参数"
// // @Success 200 {object} model.Base
// // @Failed 200 {object} model.Base
// // @Router /memo/complete [put]
// func (h *MemoHandler) CompleteWithId(ctx *gin.Context) {
// 	var request model.MemoCompleteRequest
// 	err := ctx.ShouldBind(&request)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, model.Base{
// 			Status:  model.StatusError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}
// 	err = h.usecase.CompleteWithId(ctx, request.ID)
// 	if err != nil {
// 		ctx.JSON(http.StatusOK, model.Base{
// 			Status:  model.StatusError,
// 			Message: err.Error(),
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, model.Base{
// 		Status: model.StatusOk,
// 	})
// }
