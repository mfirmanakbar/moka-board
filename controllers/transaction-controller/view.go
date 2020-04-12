package transaction_controller

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"github.com/mfirmanakbar/moka-board/utils"
)

type View struct {
	beego.Controller
}

func (c *View) Get() {
	prm := c.SearchParams()

	tms, err := models.TransactionDTO().SearchData(prm)
	if err != nil {
		println(err.Error())
		c.Data["err_msg"] = "Error on models.TransactionDTO().SearchData(prm): " + err.Error()
	}

	c.DataModifiedParams(prm)
	c.Data["tms"] = tms
	c.Data["ActiveMenu"] = "#MenuTransaction"
	c.Data["page_title"] = "Transaction Mapping"
	c.TplName = "transaction/index.html"
}

func (c *View) DataModifiedParams(prm models.TransactionParams) {
	if prm.MokaTransactionType > -1 {
		c.Data["MokaTransactionType"] = prm.MokaTransactionType
		c.Data["MokaTransactionTypeIndex"] = prm.MokaTransactionType + 1
	}
	if prm.ConnectionId > 0 {
		c.Data["ConnectionId"] = prm.ConnectionId
	}
	if prm.JurnalTransactionId > 0 {
		c.Data["JurnalTransactionId"] = prm.JurnalTransactionId
	}
	if prm.MokaTransactionId > 0 {
		c.Data["MokaTransactionId"] = prm.MokaTransactionId
	}
}

func (c *View) SearchParams() models.TransactionParams {
	params := models.TransactionParams{
		MokaTransactionType: utils.ConvertStringToInt8(c.GetString("MokaTransactionType"), -1),
		ConnectionId:        utils.ConvertStringToInt64(c.GetString("ConnectionId"), 0),
		MokaTransactionId:   utils.ConvertStringToInt64(c.GetString("MokaTransactionId"), 0),
		JurnalTransactionId: utils.ConvertStringToInt64(c.GetString("JurnalTransactionId"), 0),
	}
	return params
}
