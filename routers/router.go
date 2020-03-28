package routers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/transaction", &controllers.TransactionController{})
}
