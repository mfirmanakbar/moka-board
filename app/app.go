package app

import (
	"github.com/astaxie/beego"
	_ "github.com/mfirmanakbar/moka-board/routers" // out routers
)

func Run() {
	beego.Run()
}
