package delivery

import (
	"net/http"
	"strings"

	_ "personal_page/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func (wd *WebDeli) registerPageRouter(r *gin.Engine) {
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.GET("/load", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "load.html", nil)
	})

	r.GET("/nav", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "nav.html", nil)
	})

	r.GET("/memo", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "memo.html", nil)
	})

	r.GET("/video.html", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "video.html", nil) })

	r.GET("/movie/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if name != "" {
			ctx.HTML(http.StatusOK, "movie.html", gin.H{
				"title": strings.TrimSuffix(name, ".mp4"),
				"name":  name,
			})
		}
	})

}
