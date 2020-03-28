package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mfirmanakbar/moka-board/datasources"
	_ "github.com/mfirmanakbar/moka-board/routers"
)

func main() {
	// print out SQL queries.
	orm.Debug = true

	// init db ORM
	datasources.JurnalMokaDB()

	beego.Run()
}
