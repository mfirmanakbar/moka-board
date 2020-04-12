package connection_controller

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
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
			resultConnectionDetail := make(map[string]interface{})
			refConn := reflect.ValueOf(conn).Elem()
			for i := 0; i < refConn.NumField(); i++ {
				varName := refConn.Type().Field(i).Name
				//varType := e.Type().Field(i).Type
				varValue := refConn.Field(i).Interface()
				varName = splitUppercase(varName)
				resultConnectionDetail[varName] = varValue
			}
			c.Data["connectionDetail"] = resultConnectionDetail
		}

		jurnalUsers, _ := models.JurnalUserDTO().FindByConnectionId(connectionId)
		c.Data["jurnalUsers"] = jurnalUsers

		mokaAccount, err := models.MokaAccountDTO().FindByConnectionId(connectionId)
		if err == nil {
			resultMokaAccount := make(map[string]interface{})
			refMa := reflect.ValueOf(mokaAccount).Elem()
			for i := 0; i < refMa.NumField(); i++ {
				varName := refMa.Type().Field(i).Name
				varValue := refMa.Field(i).Interface()
				if !strings.Contains(strings.ToUpper(varName), "ACCOUNTMAPPING") {
					varName = splitUppercase(varName)
					resultMokaAccount[varName] = varValue
				}
			}
			c.Data["mokaAccount"] = resultMokaAccount
		}
	}
	// c.Data["BaseURLOldFe"] = db.BaseURLOldFe
	c.Data["ActiveMenu"] = "#MenuConnection"
	c.Data["page_title"] = "Details of Connection ID: " + connectionIdStr
	c.TplName = "connection/view.html"
}

func splitUppercase(str string) string {
	regx := regexp.MustCompile(`[A-Z][^A-Z]*`)
	words := regx.FindAllString(str, -1)
	return strings.ToUpper(strings.Join(words, " "))
}
