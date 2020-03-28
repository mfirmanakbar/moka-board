package app

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/datasources"
	_ "github.com/mfirmanakbar/moka-board/routers"
)

var (
	Db1         = datasources.JurnalMokaDbGorm{}
	MysqlDriver = "mysql"
)

func Run() {
	// print out SQL queries.
	//orm.Debug = true

	// init db ORM
	//datasources.JurnalMokaDB()

	//Db1.Initialize(
	//	MysqlDriver,
	//	os.Getenv("JURNAL_MOKA_DB_USERNAME"),
	//	os.Getenv("JURNAL_MOKA_DB_PASSWORD"),
	//	os.Getenv("JURNAL_MOKA_DB_PORT"),
	//	os.Getenv("JURNAL_MOKA_DB_HOST"),
	//	os.Getenv("JURNAL_MOKA_DB_NAME"),
	//)

	beego.Run()
}
