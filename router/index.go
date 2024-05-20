package router

import (
	"github.com/gin-gonic/gin"
	"gofly/app/handler/wx"
)

// 免accecc_token路由
func InitIndexRouter(engine *gin.Engine) {
	engine.GET("/", wx.IndexHdl.Index)
	engine.GET("/envs", wx.IndexHdl.Envs)

}
