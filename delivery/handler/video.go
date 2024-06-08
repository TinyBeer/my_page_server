package handler

import (
	"net/http"

	"personal_page/model"

	"github.com/gin-gonic/gin"
)

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
