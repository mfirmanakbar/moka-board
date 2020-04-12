package log_controller

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
)

type Index struct {
	beego.Controller
}

func (c *Index) Get() {

	logs, pgn, _ := models.LoggingSummaryDTO().FindOneByConnectionId(585)

	c.Data["paginator"] = pgn
	c.Data["loggingSummaries"] = logs
	c.Data["ActiveMenu"] = "#MenuLog"
	c.Data["page_title"] = "Log Summary"
	c.TplName = "log/index.html"
}
