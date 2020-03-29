package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/datasources"
	"github.com/mfirmanakbar/moka-board/models"
	"os"
	"strconv"
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

	println(c.GetString("mokaTransactionType"))
	println(c.GetString("connectionId"))
	println(c.GetString("mokaTransactionId"))
	println(c.GetString("jurnalTransactionId"))

	params := models.TransactionMapping{
		JurnalTransactionId: convertStringToInt64(c.GetString("jurnalTransactionId"), 0),
		MokaTransactionId:   convertStringToInt64(c.GetString("mokaTransactionId"), 0),
		MokaTransactionType: convertStringToInt8(c.GetString("mokaTransactionType"), -1),
		ConnectionId:        convertStringToInt64(c.GetString("connectionId"), 0),
	}

	tm := models.TransactionMapping{}
	tms, err := tm.SearchTransactionMappings(conn, params.ConnectionId)
	if err != nil {
		tms = nil
	}

	println(params.MokaTransactionType)
	println(params.ConnectionId)
	println(params.MokaTransactionId)
	println(params.JurnalTransactionId)

	if params.MokaTransactionType > -1 {
		c.Data["MokaTransactionType"] = params.MokaTransactionType
	}

	if params.ConnectionId > 0 {
		c.Data["ConnectionId"] = params.ConnectionId
	}

	if params.JurnalTransactionId > 0 {
		c.Data["JurnalTransactionId"] = params.JurnalTransactionId
	}

	if params.MokaTransactionId > 0 {
		c.Data["MokaTransactionId"] = params.MokaTransactionId
	}

	c.Data["params"] = params
	c.Data["tms"] = tms
	c.Data["page_title"] = "Transaction Mapping"
	c.TplName = "transaction/index.html"

	//	//itemx := []int{1, 2, 3, 4, 5}
	//	//c.Data["itemx"] = itemx

	//	//services.FindTransactionMappings(585, 242590, 1284563, 1)
	//	//services.FindTransactionMappings(585, 242590, 0, -1)
}

func convertStringToInt8(str string, defaultResult int) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return defaultResult
	}
	return result
}

func convertStringToInt64(str string, defaultResult int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultResult
	}
	return result
}
