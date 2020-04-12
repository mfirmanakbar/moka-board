package connection_controller

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"github.com/mfirmanakbar/moka-board/utils"
)

type Index struct {
	beego.Controller
}

func (c *Index) Get() {
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

func (c *Index) DataModifiedParams(prm models.ConnectionParams) {
	//searchResultsOf := map[string]string{"ab1":"cd", "ab2":"go"}
	//searchResultsOf := []string{"abc","def"}
	searchResultsOf := make(map[string]interface{})

	if prm.CompanyId > 0 {
		c.Data["CompanyId"] = prm.CompanyId
		searchResultsOf["CompanyID"] = prm.CompanyId
	}

	if prm.ConnectionId > 0 {
		c.Data["ConnectionId"] = prm.ConnectionId
		searchResultsOf["ConnectionID"] = prm.ConnectionId
	}

	if prm.ConnectionName != "" {
		searchResultsOf["ConnectionName"] = prm.ConnectionName
	}

	if prm.ShowDeleted {
		searchResultsOf["ShowDeleted"] = "Yes"
	}

	if prm.ShowOnlySyncing {
		searchResultsOf["ShowOnlySyncing"] = "Yes"
	}

	c.Data["searchResultsOf"] = searchResultsOf
	c.Data["ConnectionName"] = prm.ConnectionName
	c.Data["ShowDeleted"] = prm.ShowDeleted
	c.Data["ShowOnlySyncing"] = prm.ShowOnlySyncing
}

func (c *Index) SearchParams() models.ConnectionParams {
	params := models.ConnectionParams{
		CompanyId:       utils.ConvertStringToInt64(c.GetString("CompanyId"), 0),
		ConnectionId:    utils.ConvertStringToInt64(c.GetString("ConnectionId"), 0),
		ConnectionName:  c.GetString("ConnectionName"),
		ShowDeleted:     utils.ConvertStringToBool(c.GetString("ShowDeleted")),
		ShowOnlySyncing: utils.ConvertStringToBool(c.GetString("ShowOnlySyncing")),
	}
	return params
}
