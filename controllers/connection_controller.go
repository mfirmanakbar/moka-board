package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"github.com/mfirmanakbar/moka-board/utils"
)

var (
	cs = models.Connection{}
)

type ConnectionController struct {
	beego.Controller
}

type ConnectionControllerInterface interface {
	SearchParams() models.ConnectionParams
	DataModifiedParams(models.ConnectionParams)
}

func Connection() ConnectionControllerInterface {
	return &ConnectionController{}
}

func (c *ConnectionController) Get() {
	prm := c.SearchParams()

	conns, err := models.ConnectionDTO().SearchData(prm)
	if err != nil {
		println(err.Error())
		c.Data["err_msg"] = "Error on models.ConnectionDTO().SearchData(prm): " + err.Error()
	}

	c.DataModifiedParams(prm)
	c.Data["conns"] = conns
	c.Data["ActiveMenu"] = "#MenuConnection"
	c.Data["page_title"] = "Connection Settings"
	c.TplName = "connection/index.html"
}

func (c *ConnectionController) DataModifiedParams(prm models.ConnectionParams) {
	if prm.CompanyId > 0 {
		c.Data["CompanyId"] = prm.CompanyId
	}
	if prm.ConnectionId > 0 {
		c.Data["ConnectionId"] = prm.ConnectionId
	}
	c.Data["ConnectionName"] = prm.ConnectionName
	c.Data["ShowDeleted"] = prm.ShowDeleted
	c.Data["ShowOnlySyncing"] = prm.ShowOnlySyncing
}

func (c *ConnectionController) SearchParams() models.ConnectionParams {
	params := models.ConnectionParams{
		CompanyId:       utils.ConvertStringToInt64(c.GetString("CompanyId"), 0),
		ConnectionId:    utils.ConvertStringToInt64(c.GetString("ConnectionId"), 0),
		ConnectionName:  c.GetString("ConnectionName"),
		ShowDeleted:     utils.ConvertStringToBool(c.GetString("ShowDeleted")),
		ShowOnlySyncing: utils.ConvertStringToBool(c.GetString("ShowOnlySyncing")),
	}
	return params
}
