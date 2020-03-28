package controllers

import (
	"github.com/astaxie/beego"
)

type TransactionController struct {
	beego.Controller
}

func (c *TransactionController) Get() {

	println(c.GetString("mokaTransactionType"))
	println(c.GetString("connectionId"))
	println(c.GetString("mokaTransactionId"))
	println(c.GetString("jurnalTransactionId"))

	c.TplName = "transaction/index.html"
	c.Data["page_title"] = "Transaction Mapping"

	//itemx := []int{1, 2, 3, 4, 5}
	//c.Data["itemx"] = itemx

	//services.FindTransactionMappings(585, 242590, 1284563, 1)
	//services.FindTransactionMappings(585, 242590, 0, -1)
}
