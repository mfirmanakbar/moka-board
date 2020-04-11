package connection_controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"reflect"
	"strconv"
)

type View struct {
	beego.Controller
}

func (c *View) Get() {
	connectionIdStr := c.Ctx.Input.Param(":id")
	if connectionIdStr != "" {
		connectionId, _ := strconv.ParseInt(connectionIdStr, 10, 64)
		prm := models.ConnectionParams{
			ConnectionId: connectionId,
			ShowDeleted:  true,
		}
		conn, err := models.ConnectionDTO().FindById(prm)
		if err == nil {
			resultOfReflect := make(map[string]interface{})
			e := reflect.ValueOf(conn).Elem()
			for i := 0; i < e.NumField(); i++ {
				varName := e.Type().Field(i).Name
				varType := e.Type().Field(i).Type
				varValue := e.Field(i).Interface()
				fmt.Printf("%v %v %v\n", varName, varType, varValue)
				resultOfReflect[varName] = varValue
			}
			c.Data["connectionDetail"] = resultOfReflect
		}
	}
	c.Data["ActiveMenu"] = "#MenuConnection"
	c.Data["page_title"] = "Details of Connection ID: " + connectionIdStr
	c.TplName = "connection/view.html"
}
