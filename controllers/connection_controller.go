package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"github.com/mfirmanakbar/moka-board/utils"
)

type ConnectionController struct {
	beego.Controller
}

func (c *ConnectionController) get() {
	prm := c.SearchParams()

	if prm.CompanyId > 0 {
		conns, err := models.JurnalCompanyDTO().FindConnectionsByCompanyId2(prm.CompanyId, prm.ShowDeleted)
		if err != nil {
			println(err.Error())
		}
		c.Data["conns"] = conns
	} else {
		conns, err := models.ConnectionDTO().SearchData(prm)
		if err != nil {
			println(err.Error())
			c.Data["err_msg"] = "Error on models.ConnectionDTO().SearchData(prm): " + err.Error()
		}
		c.Data["conns"] = conns
	}

	c.DataModifiedParams(prm)
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
