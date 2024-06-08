package delivery

import (
	"io"
	"net/http"

	"personal_page/usecase"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type WebDeli struct {
	uc  usecase.UserUsecase
	muc usecase.MemoUsecase
	out io.Writer
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
	pprof.Register(r)
	return r
}

// Start implements WebDelivery.
func (w *WebDeli) Start() {
	r := w.router()
	r.Run(":9999")
}

func NewWebDeli(uc usecase.UserUsecase, muc usecase.MemoUsecase) *WebDeli {
	return &WebDeli{
		uc:  uc,
		muc: muc,
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
