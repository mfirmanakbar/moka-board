package connection_controller

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"time"
	"fmt"
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

		// loc, err := time.LoadLocation("America/Los_Angeles")
    // if err != nil {
    //     fmt.Println(err)
    // }
		// fmt.Println("Firman")
		// fmt.Println(loc)
		
		conn, err := models.ConnectionDTO().FindById(prm)

		fmt.Println("firman")
		fmt.Println(conn.StartDate)

		// fmt.Println(conn.StartDate.Format("2006-01-02 15:04:05"))
		//jktTime, _ := time.LoadLocation("Asia/Singapore")
		//fmt.Println(conn.StartDate.In(jktTime))
		// fmt.Println(conn.StartDate.UTC())
		loc, _ := time.LoadLocation("Asia/Jakarta")
		//now := time.Now().In(loc)
		now := conn.StartDate.In(loc)
    fmt.Println("Location : ", loc, " Time : ", now)


		if err == nil {
			c.Data["connDetail"] = conn
		}

	}
	c.Data["ActiveMenu"] = "#MenuConnection"
	c.Data["page_title"] = "Details of Connection ID: " + connectionIdStr
	c.TplName = "connection/edit.html"
}
