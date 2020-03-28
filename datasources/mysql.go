package datasources

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func JurnalMokaDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "useronlyedit:P@ssw0rd@/jurnal_moka_db?charset=utf8")
}
