package delivery

import (
	"personal_page/delivery/handler"
	"personal_page/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func (wd *WebDeli) registerUserRouter(r *gin.Engine) {
	uh := handler.NewUserHandler(wd.uc)

	user := r.Group("/user")
	user.POST("/login", uh.Login)
	user.POST("/auth", uh.Auth)
	user.POST("/refresh", uh.RefreshToken)
}

func (wd *WebDeli) registerVideoRouter(r *gin.Engine) {
	md := middleware.NewMiddleware(wd.uc)
	video := r.Group("/video", md.JWT)
	{
		video.GET("/list", handler.VideoHandler{}.List)
	}
}

func (wd *WebDeli) registerMemoRouter(r *gin.Engine) {
	uh := handler.NewMemoHandler(wd.muc)
	md := middleware.NewMiddleware(wd.uc)
	memo := r.Group("/memo", md.JWT)
	{
		memo.GET("/list", uh.List)
		memo.POST("/create", uh.Create)
		memo.DELETE("/delete", uh.DeleteById)
		memo.PUT("/update", uh.UpdateById)
	}
}
