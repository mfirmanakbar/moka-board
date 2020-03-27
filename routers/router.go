package routers

import (
	"github.com/astaxie/beego"
	"moka-board/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
