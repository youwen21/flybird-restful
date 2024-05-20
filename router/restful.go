package router

import (
	"github.com/gin-gonic/gin"
	"gofly/app/rest/handler"
)

// 免accecc_token路由
func InitRestfulRouter(engine *gin.Engine) {
	engine.GET("/restful/:table", handler.RestHdl.Query)
	engine.GET("/restful/:table/:id", handler.RestHdl.Get)

	engine.PUT("/restful/:table", handler.RestHdl.Insert)
	engine.POST("/restful/:table/:id", handler.RestHdl.Update)
	engine.DELETE("/restful/:table/:id", handler.RestHdl.Delete)

	engine.POST("/access_db/query", handler.SqlHdl.Query)

	engine.POST("/access_db/insert", handler.SqlHdl.Insert)
	engine.POST("/access_db/update", handler.SqlHdl.Update)
	engine.POST("/access_db/delete", handler.SqlHdl.Delete)

	engine.POST("/access_db/sqlQuery", handler.SqlHdl.SqlRawQuery)
	engine.POST("/access_db/sqlExecute", handler.SqlHdl.Execute)
}
