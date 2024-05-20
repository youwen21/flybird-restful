package handler

import (
	"github.com/gin-gonic/gin"
	"gofly/app/rest/forms"
	"gofly/app/rest/service"
	"gofly/gin_helper"
)

type sqlHdl struct {
}

var SqlHdl = &sqlHdl{}

func (h *sqlHdl) Query(c *gin.Context) {
	form := new(forms.QueryForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	res, err := service.SqlSrv.Query(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	gin_helper.Json(c, res, nil)
}

func (h *sqlHdl) SqlRawQuery(c *gin.Context) {
	form := new(forms.SqlQueryForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	res, err := service.SqlSrv.SqlRawQuery(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	gin_helper.Json(c, gin.H{"columns": nil, "res": res}, nil)
}

func (h *sqlHdl) Get(c *gin.Context) {
	form := new(forms.GetForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	res, err := service.SqlSrv.Get(form)
	gin_helper.Json(c, res, err)
}

func (h *sqlHdl) Insert(c *gin.Context) {
	form := new(forms.PutForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	res, err := service.SqlSrv.Insert(form)
	gin_helper.Json(c, res, err)
}

func (h *sqlHdl) Update(c *gin.Context) {
	form := new(forms.RestUpdateForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	res, err := service.SqlSrv.Update(form)
	gin_helper.Json(c, res, err)
}

func (h *sqlHdl) Delete(c *gin.Context) {
	form := new(forms.GetForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	res, err := service.SqlSrv.Delete(form)
	gin_helper.Json(c, res, err)
}

func (h *sqlHdl) Execute(c *gin.Context) {

	form := new(forms.SqlExecuteForm)
	err := c.ShouldBind(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}

	err = service.SqlSrv.Execute(form)
	if err != nil {
		gin_helper.JsonErr(c, err)
		return
	}
	gin_helper.Json(c, nil, nil)
}
