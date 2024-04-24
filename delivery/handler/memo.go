package handler

import (
	"net/http"
	"personal_page/model"
	"personal_page/usecase"

	"github.com/gin-gonic/gin"
)

type MemoHandler struct {
	usecase usecase.MemoUsecase
}

func NewMemoHandler(uc usecase.MemoUsecase) *MemoHandler {
	return &MemoHandler{
		usecase: uc,
	}
}

// List 获取备忘录列表接口
// @Summary 获取备忘录列表资源
// @Description 使用access_token获取视频列表资源
// @Tags 备忘录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} model.MemoListResponse
// @Router /memo/list [get]
func (h *MemoHandler) List(ctx *gin.Context) {
	list, err := h.usecase.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, model.MemoListResponse{
			Status:  "error",
			Message: err.Error(),
		})
	} else {
		memoes := make([]model.MemoItem, 0, len(list))
		for _, v := range list {
			memoes = append(memoes, v.ToMemoItem())
		}
		ctx.JSON(http.StatusOK, model.MemoListResponse{
			Status:  "success",
			Message: "",
			Memoes:  memoes,
		})
	}
}

// Create 创建备忘录
// @Summary 根据内容创建新的备忘录
// @Description 使用access_token获取视频列表资源
// @Tags 备忘录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body model.MemoCreateRequest true "备忘录创建参数"
// @Success 200 {object} model.SimpleResponse
// @Router /memo/create [post]
func (h *MemoHandler) Create(ctx *gin.Context) {
	var request model.MemoCreateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusOK, model.SimpleResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	err = h.usecase.Create(ctx, request.Content)
	if err != nil {
		ctx.JSON(http.StatusOK, model.SimpleResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.SimpleResponse{
		Status:  "success",
		Message: "",
	})
}

// DeleteById 删除备忘录
// @Summary 根据Id删除备忘录
// @Description 使用access_token获取视频列表资源
// @Tags 备忘录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body model.MemoDeleteRequest true "备忘录删除参数"
// @Success 200 {object} model.SimpleResponse
// @Failed 200 {object} model.SimpleResponse
// @Router /memo/delete [delete]
func (h *MemoHandler) DeleteById(ctx *gin.Context) {
	var request model.MemoDeleteRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusOK, model.SimpleResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	err = h.usecase.DeleteById(ctx, request.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, model.SimpleResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.SimpleResponse{
		Status: "success",
	})
}

// UpdateById 更新备忘录
// @Summary 根据Id更新备忘录
// @Description 使用access_token获取视频列表资源
// @Tags 备忘录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body model.MemoUpdateRequest true "备忘录删除参数"
// @Success 200 {object} model.SimpleResponse
// @Failed 200 {object} model.SimpleResponse
// @Router /memo/update [put]
func (h *MemoHandler) UpdateById(ctx *gin.Context) {
	var request model.MemoUpdateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusOK, model.SimpleResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	err = h.usecase.UpdateById(ctx, request.Id, request.Content)
	if err != nil {
		ctx.JSON(http.StatusOK, model.SimpleResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.SimpleResponse{
		Status: "success",
	})
}
