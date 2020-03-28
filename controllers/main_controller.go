package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "mfirmanakbar.com"
	c.Data["Email"] = "em.firman.akbar@gmail.com"
	c.TplName = "index.html"
}
