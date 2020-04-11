package routers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/controllers"
	"github.com/mfirmanakbar/moka-board/controllers/connection_controller"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/transaction", &controllers.TransactionController{})
	beego.Router("/connection", &connection_controller.Index{})
	beego.Router("/connection/:id([0-9]+/view", &connection_controller.View{})
}
