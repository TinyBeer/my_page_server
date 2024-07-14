package delivery

import (
	"fmt"

	"personal_page/domain"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type WebDeli struct {
	port  string
	token domain.TokenUsecase
	todo  domain.TodoUsecase
	plan  domain.PlanUsecase
}

func (wd *WebDeli) router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), Cors())

	wd.registerSupportPage(r)
	wd.registerTokenRouter(r)
	wd.registerTodoRouter(r)
	wd.registerPlanRouter(r)
	return r
}

// Start implements WebDelivery.
func (w *WebDeli) Start() {
	r := w.router()
	r.Run(w.port)
}

func NewWebDeli(
	conf *viper.Viper,
	token domain.TokenUsecase,
	todo domain.TodoUsecase,
	plan domain.PlanUsecase,
) *WebDeli {
	return &WebDeli{
		port:  fmt.Sprintf(":%d", conf.GetInt("server.port")),
		token: token,
		todo:  todo,
		plan:  plan,
	}
}

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
