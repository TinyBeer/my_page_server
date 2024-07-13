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
	user.POST("/refresh", uh.RefreshToken)

	md := middleware.NewMiddleware(wd.uc)
	user.POST("/auth", md.JWT, uh.Auth)
}

func (wd *WebDeli) registerVideoRouter(r *gin.Engine) {
	md := middleware.NewMiddleware(wd.uc)
	video := r.Group("/video", md.JWT)
	{
		video.GET("/list", handler.VideoHandler{}.List)
	}
}

func (wd *WebDeli) registerMovieRouter(r *gin.Engine) {
	md := middleware.NewMiddleware(wd.uc)
	mvh := handler.NewMovieHandler(wd.mvuc)
	video := r.Group("/movie", md.JWT)
	{
		video.GET("/list", mvh.List)
		video.POST("/create", mvh.Create)
		video.DELETE("/delete", mvh.DeleteById)
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
		memo.PUT("/complete", uh.CompleteWithId)
	}
}

func (wd *WebDeli) registerTodoRouter(r *gin.Engine) {
	h := handler.NewTodoHandler(wd.todo)
	md := middleware.NewMiddleware(wd.uc)
	todo := r.Group("/todo", md.JWT)
	{
		todo.GET("", h.List)
		todo.POST("", h.Create)
		todo.DELETE("", h.DeleteByID)
	}
}
