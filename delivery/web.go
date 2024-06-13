package delivery

import (
	"fmt"
	"io"
	"net/http"

	"personal_page/usecase"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type WebDeli struct {
	port string
	uc   usecase.UserUsecase
	muc  usecase.MemoUsecase
	mvuc usecase.MovieUsecase
	out  io.Writer
}

func (wd *WebDeli) router() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(wd.out), gin.Recovery(), Cors())

	r.Static("/static", "./static")

	r.LoadHTMLGlob("./html/*")

	r.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/nav")
	})

	wd.registerPageRouter(r)
	wd.registerUserRouter(r)
	wd.registerVideoRouter(r)
	wd.registerMemoRouter(r)
	wd.registerMovieRouter(r)
	pprof.Register(r)
	return r
}

// Start implements WebDelivery.
func (w *WebDeli) Start() {
	r := w.router()
	r.Run(w.port)
}

func NewWebDeli(conf *viper.Viper, uc usecase.UserUsecase, muc usecase.MemoUsecase, mvuc usecase.MovieUsecase) *WebDeli {
	return &WebDeli{
		port: fmt.Sprintf(":%d", conf.GetInt("server.port")),
		uc:   uc,
		muc:  muc,
		mvuc: mvuc,
		out:  nil,
	}
}

var _ WebDelivery = &WebDeli{}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}
