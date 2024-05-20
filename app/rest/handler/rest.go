package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gofly/app/rest/forms"
	"gofly/app/rest/service"
	"gofly/gin_helper"
)

type restHdl struct {
}

var RestHdl = &restHdl{}

func (h *restHdl) Query(c *gin.Context) {
	form := new(forms.QueryForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}
	form.Table = c.Param("table")

	ret, err := service.SqlSrv.Query(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}
	gin_helper.Json(c, ret, err)
}

func (h *restHdl) Get(c *gin.Context) {
	form := new(forms.GetForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}
	form.Table = c.Param("table")
	form.Id = cast.ToInt(c.Param("id"))

	res, err := service.SqlSrv.Get(form)
	gin_helper.Json(c, res, err)
}

func (h *restHdl) Insert(c *gin.Context) {
	form := new(forms.PutForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}
	form.Table = c.Param("table")

	res, err := service.SqlSrv.Insert(form)
	gin_helper.Json(c, res, err)
}

func (h *restHdl) Update(c *gin.Context) {
	form := new(forms.RestUpdateForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}
	form.Table = c.Param("table")
	form.Id = cast.ToInt(c.Param("id"))

	res, err := service.SqlSrv.Update(form)
	gin_helper.Json(c, res, err)
}

func (h *restHdl) Delete(c *gin.Context) {
	form := new(forms.GetForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}
	form.Table = c.Param("table")
	form.Id = cast.ToInt(c.Param("id"))

	res, err := service.SqlSrv.Delete(form)
	gin_helper.Json(c, res, err)
}
