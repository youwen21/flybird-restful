package gin_helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofly/apperror"
	"net/http"
)

func Json(c *gin.Context, data interface{}, err error) {
	code := 200
	msg := "success"

	if err != nil {
		code, msg = GetErrorCodeMsg(err)
	}

	JsonRaw(c, code, msg, data)
}

func JsonErr(c *gin.Context, e error) {
	JsonRaw(c, getErrorCode(e), e.Error(), nil)
}

func JsonRaw(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
}

func getErrorCode(e error) int {
	if appErr, ok := e.(apperror.AppError); ok {
		return appErr.Code
	}
	return 9999
}

func GetErrorCodeMsg(e error) (int, string) {
	if e == nil {
		return 0, ""
	}

	if appErr, ok := e.(apperror.AppError); ok {
		return appErr.Code, appErr.Error()
	}
	return 9999, e.Error()
}

func Refresh(c *gin.Context, timeout int, uri string, title string, msg string) {
	//c.String(http.StatusOK, pageRefresh, timeout, uri, title, msg, timeout)
	pageHtml := fmt.Sprintf(pageRefresh, timeout, uri, title, msg, timeout)
	c.Writer.Write([]byte(pageHtml))
	//c.String(http.StatusOK, pageRefresh, timeout, uri, title, msg, timeout)
}
