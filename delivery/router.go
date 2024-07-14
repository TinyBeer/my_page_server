package delivery

import (
	"personal_page/delivery/handler"
	"personal_page/delivery/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	_ "personal_page/docs"
)

func (wd *WebDeli) registerTokenRouter(r *gin.Engine) {
	h := handler.NewTokenHandler(wd.token)
	token := r.Group("/token")
	token.POST("", h.Gen)
	token.POST("/access_token", h.AccessToken)

	md := middleware.NewMiddleware(wd.token)
	token.GET("/access_token", md.JWT, h.Check)
}

func (wd *WebDeli) registerTodoRouter(r *gin.Engine) {
	h := handler.NewTodoHandler(wd.todo)
	md := middleware.NewMiddleware(wd.token)
	todo := r.Group("/todo", md.JWT)
	{
		todo.GET("", h.List)
		todo.POST("", h.Create)
		todo.DELETE("", h.DeleteByID)
		todo.PUT("", h.CompleteByID)
	}
}

func (wd *WebDeli) registerSupportPage(r *gin.Engine) {
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	pprof.Register(r)
}
