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
	r.Use(gin.LoggerWithWriter(wd.out), gin.Recovery())

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

func NewWebDeli(uc usecase.UserUsecase, muc usecase.MemoUsecase, out io.Writer) *WebDeli {
	return &WebDeli{
		uc:  uc,
		muc: muc,
		out: out,
	}
}

var _ WebDelivery = &WebDeli{}
