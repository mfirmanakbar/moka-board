package routers

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/controllers"
	ConnectionController "github.com/mfirmanakbar/moka-board/controllers/connection-controller"
	LogConnection "github.com/mfirmanakbar/moka-board/controllers/log-controller"
	TransactionConnection "github.com/mfirmanakbar/moka-board/controllers/transaction-controller"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/transaction", &TransactionConnection.View{})
	beego.Router("/connection", &ConnectionController.Index{})
	beego.Router("/connection/:id([0-9]+/view", &ConnectionController.View{})
	beego.Router("/connection/:id([0-9]+/edit", &ConnectionController.Edit{})
	beego.Router("/log", &LogConnection.Index{})
}
