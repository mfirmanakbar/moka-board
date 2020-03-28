package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Website"] = "mfirmanakbar.com"
	c.Data["Email"] = "em.firman.akbar@gmail.com"
	c.TplName = "login/index.html"
}
