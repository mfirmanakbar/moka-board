package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"github.com/mfirmanakbar/moka-board/utils"
)

var (
	tm = models.TransactionMapping{}
)

type TransactionController struct {
	beego.Controller
}

func (c *TransactionController) Get() {
	prm := SearchParams(c)

	tms, err := tm.SearchTransactionMappings(prm)
	if err != nil {
		println(err.Error())
		c.Data["err_msg"] = "Error on SearchTransactionMappings: " + err.Error()
	}

	DataModifiedParams(prm, c)

	c.Data["ActiveMenu"] = "#MenuTransaction"
	c.Data["tms"] = tms
	c.Data["page_title"] = "Transaction Mapping"
	c.TplName = "transaction/index.html"
}

func DataModifiedParams(tm models.SearchParams, c *TransactionController) {
	if tm.MokaTransactionType > -1 {
		c.Data["MokaTransactionType"] = tm.MokaTransactionType
		c.Data["MokaTransactionTypeIndex"] = tm.MokaTransactionType + 1
	}

	if tm.ConnectionId > 0 {
		c.Data["ConnectionId"] = tm.ConnectionId
	}

	if tm.JurnalTransactionId > 0 {
		c.Data["JurnalTransactionId"] = tm.JurnalTransactionId
	}

	if tm.MokaTransactionId > 0 {
		c.Data["MokaTransactionId"] = tm.MokaTransactionId
	}
}

func SearchParams(c *TransactionController) models.SearchParams {
	params := models.SearchParams{
		MokaTransactionType: utils.ConvertStringToInt8(c.GetString("MokaTransactionType"), -1),
		ConnectionId:        utils.ConvertStringToInt64(c.GetString("ConnectionId"), 0),
		MokaTransactionId:   utils.ConvertStringToInt64(c.GetString("MokaTransactionId"), 0),
		JurnalTransactionId: utils.ConvertStringToInt64(c.GetString("JurnalTransactionId"), 0),
	}
	return params
}
