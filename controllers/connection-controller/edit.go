package connection_controller

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
)

type Edit struct {
	beego.Controller
}

func (c *Edit) Get() {
	connectionIdStr := c.Ctx.Input.Param(":id")
	if connectionIdStr != "" {
		connectionId, _ := strconv.ParseInt(connectionIdStr, 10, 64)
		prm := models.ConnectionParams{
			ConnectionId: connectionId,
			ShowDeleted:  true,
		}

		conn, err := models.ConnectionDTO().FindById(prm)
		if err == nil {
			c.Data["connDetail"] = conn
		}

	}
	c.Data["ActiveMenu"] = "#MenuConnection"
	c.Data["page_title"] = "Details of Connection ID: " + connectionIdStr
	c.TplName = "connection/edit.html"
}
