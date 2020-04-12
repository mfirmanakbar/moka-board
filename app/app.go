package app

import (
	"github.com/astaxie/beego"
	_ "github.com/mfirmanakbar/moka-board/routers" // out routers
)

// Run the router
func Run() {
	beego.Run()
}
