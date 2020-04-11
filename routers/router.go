package routers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/transaction", &controllers.TransactionController{})
	beego.Router("/connection", &controllers.ConnectionController{})
	//beego.Router("/connection/:id/view", &controllers.ConnectionController{})
}
