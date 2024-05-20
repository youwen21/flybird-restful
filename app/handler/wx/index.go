package wx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofly/gin_helper"
	"net/http"
	"os"
)

/*  */

type indexHdl struct{}

var (
	IndexHdl = &indexHdl{}
)

func (cHdl *indexHdl) Index(c *gin.Context) {
	content := `Hello, Welcome to FlyBird Restful<br/>
`
	c.String(http.StatusOK, content)

}

func (cHdl *indexHdl) Envs(c *gin.Context) {
	envs := os.Environ()
	fmt.Println(envs)
	gin_helper.Json(c, envs, nil)
}
