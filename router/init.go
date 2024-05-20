package router

import (
	"github.com/gin-gonic/gin"
	"gofly/middleware"
)

func InitRouter(engine *gin.Engine) {
	//gin.SetMode(gin.DebugMode)
	engine.GET("/ping", func(gtx *gin.Context) { gtx.String(200, "pong") })
	engine.OPTIONS("/*options_support", middleware.Cors.GinCors())

	InitIndexRouter(engine)

	InitRestfulRouter(engine)
}
