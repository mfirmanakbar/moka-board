package log_controller

import (
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"github.com/mfirmanakbar/moka-board/utils"
)

type Index struct {
	beego.Controller
}

func (c *Index) Get() {

	ConnectionId := utils.ConvertStringToInt64(c.GetString("ConnectionId"), 0)
	Page := utils.ConvertStringToInt8(c.GetString("Page"), 0)
	if Page == 0 {
		Page = 1
	}

	logs, pgn, _ := models.LoggingSummaryDTO().FindOneByConnectionId(ConnectionId, Page)

	c.Data["pgn"] = pgn
	c.Data["loggingSummaries"] = logs
	c.Data["ConnectionId"] = ConnectionId
	//c.Ctx.Request.Host + c.Ctx.Request.URL.Path + "?" + c.Ctx.Request.URL.RawQuery
	RawQuery := "?ConnectionId=" + c.GetString("ConnectionId") + "&Page="
	c.Data["UrlPath"] = c.Ctx.Request.Host + c.Ctx.Request.URL.Path + RawQuery
	c.Data["ActiveMenu"] = "#MenuLog"
	c.Data["page_title"] = "Log Summary"
	c.TplName = "log/index.html"
}
