package handler

import (
	"net/http"

	"personal_page/model"
	"personal_page/usecase"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	usecase usecase.MovieUsecase
}

func NewMovieHandler(uc usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{
		usecase: uc,
	}
}

// List 获取电影列表接口
// @Summary 获取电影列表资源
// @Description 使用access_token获取视频列表资源
// @Tags 电影相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} model.MovieListResponse
// @Router /movie/list [get]
func (h *MovieHandler) List(ctx *gin.Context) {
	list, err := h.usecase.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Base{
			Status:  model.StatusError,
			Message: err.Error(),
		})
	} else {
		movies := make([]model.MovieItem, 0, len(list))
		for _, v := range list {
			movies = append(movies, v.ToMovieItem())
		}
		ctx.JSON(http.StatusOK, model.MovieListResponse{
			Base: model.Base{
				Status:  model.StatusOk,
				Message: "",
			},
			Movies: movies,
		})
	}
}

// Create 创建电影资源
// @Summary 根据内容创建新的电影资源
// @Description 使用access_token创建新的电影资源
// @Tags 电影相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body model.MovieCreateRequest true "电影创建参数"
// @Success 200 {object} model.Base
// @Router /movie/create [post]
func (h *MovieHandler) Create(ctx *gin.Context) {
	var request model.MovieCreateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Base{
			Status:  model.StatusError,
			Message: err.Error(),
		})
		return
	}

	movie := &model.Movie{
		ID:     "",
		Post:   request.Post,
		Title:  request.Title,
		Tags:   request.Tags,
		Desc:   request.Desc,
		Source: request.Source,
	}
	err = h.usecase.Create(ctx, movie)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Base{
			Status:  model.StatusError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Base{
		Status:  model.StatusOk,
		Message: "",
	})
}

// DeleteById 删除电影资源
// @Summary 根据Id删除电影
// @Description 使用access_token删除电影资源
// @Tags 电影相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Param {object} body model.MovieDeleteRequest true "电影删除参数"
// @Success 200 {object} model.Base
// @Failed 200 {object} model.Base
// @Router /movie/delete [delete]
func (h *MovieHandler) DeleteById(ctx *gin.Context) {
	var request model.MovieDeleteRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Base{
			Status:  model.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.usecase.DeleteById(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Base{
			Status:  model.StatusError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Base{
		Status: model.StatusOk,
	})
}

type VideoHandler struct{}

// List 获取视频列表接口
// @Summary 获取视频列表资源
// @Description 使用access_token获取视频列表资源
// @Tags 视频相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "访问令牌"
// @Success 200 {object} model.VideoListResponse
// @Router /video [get]
func (h VideoHandler) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.listResponse())
}

func (h VideoHandler) listResponse() *model.VideoListResponse {
	return &model.VideoListResponse{
		Base: model.Base{
			Status:  model.StatusOk,
			Message: "",
		},
		Videoes: []model.VideoItem{
			{
				Image: "fastx.jpg",
				Title: "fastx",
				Time:  "2022-03-09",
				Intro: "introduction",
				Name:  "fastx.mp4",
			},
			{
				Image: "god_father.jpg",
				Title: "god_father",
				Time:  "2022-04-09",
				Intro: "introduction",
				Name:  "god_father.mp4",
			},
			{
				Image: "legend.jpg",
				Title: "legend",
				Time:  "2022-03-19",
				Intro: "introduction",
				Name:  "legend.mp4",
			},
			{
				Image: "lion.jpg",
				Title: "lion",
				Time:  "2022-03-26",
				Intro: "introduction",
				Name:  "lion.mp4",
			},
			{
				Image: "old_fox.jpg",
				Title: "old_fox",
				Time:  "2022-04-09",
				Intro: "introduction",
				Name:  "old_fox.mp4",
			},
			{
				Image: "pianist.jpg",
				Title: "pianist",
				Time:  "2022-04-09",
				Intro: "introduction",
				Name:  "pianist.mp4",
			},
		},
	}
}
