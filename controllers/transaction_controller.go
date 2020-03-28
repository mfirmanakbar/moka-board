package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/datasources"
	"github.com/mfirmanakbar/moka-board/models"
	"os"
)

var (
	Db1         = datasources.JurnalMokaDbGorm{}
	MysqlDriver = "mysql"
)

var (
	server  = datasources.JurnalMokaDbGorm{}
	conn, _ = datasources.Initialize(
		MysqlDriver,
		os.Getenv("JURNAL_MOKA_DB_USERNAME"),
		os.Getenv("JURNAL_MOKA_DB_PASSWORD"),
		os.Getenv("JURNAL_MOKA_DB_PORT"),
		os.Getenv("JURNAL_MOKA_DB_HOST"),
		os.Getenv("JURNAL_MOKA_DB_NAME"),
	)
)

type TransactionController struct {
	beego.Controller
}

func (c *TransactionController) Get() {

	// user := models.User{}
	//users, err := user.FindAllUsers(server.DB)

	/*
	   OPEN_SALES_INVOICE(0), //
	   PAID_SALES_INVOICE(1), // Hijau
	   RECEIVE_PAYMENT(2), //
	   SALES_RETURN(3), //
	   REFUNDED_OPEN_SALES_INVOICE(4); //
	*/

	tm := models.TransactionMapping{}
	tms, err := tm.SearchTransactionMappings(conn, 585)
	if err != nil {
		tms = nil
	}

	c.Data["tms"] = tms
	c.Data["page_title"] = "Transaction Mapping"
	c.TplName = "transaction/index.html"

	//	println(c.GetString("mokaTransactionType"))
	//	println(c.GetString("connectionId"))
	//	println(c.GetString("mokaTransactionId"))
	//	println(c.GetString("jurnalTransactionId"))

	//	//itemx := []int{1, 2, 3, 4, 5}
	//	//c.Data["itemx"] = itemx

	//	//services.FindTransactionMappings(585, 242590, 1284563, 1)
	//	//services.FindTransactionMappings(585, 242590, 0, -1)
}
