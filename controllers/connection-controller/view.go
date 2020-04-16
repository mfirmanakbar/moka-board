package connection_controller

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/mfirmanakbar/moka-board/models"
	"github.com/mfirmanakbar/moka-board/myenv"
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
			c.Data["connDetail"] = conn
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

		jurnalCompany, _ := models.ConnectionDTO().FindJurnalCompanyByConnId(connectionId)
		if err == nil {
			resultJurnalCompany := make(map[string]interface{})
			refMa := reflect.ValueOf(jurnalCompany).Elem()
			for i := 0; i < refMa.NumField(); i++ {
				varName := refMa.Type().Field(i).Name
				varValue := refMa.Field(i).Interface()
				if !strings.Contains(strings.ToUpper(varName), "ACCOUNTMAPPING") && !strings.Contains(strings.ToUpper(varName), "JURNALUSERS") {
					varName = splitUppercase(varName)
					resultJurnalCompany[varName] = varValue
				}
			}
			c.Data["jurnalCompany"] = resultJurnalCompany
		}

	}
	c.Data["BaseURLOldFe"] = myenv.BaseURLOldFe
	c.Data["BaseURLNewFe"] = myenv.BaseURLNewFe
	c.Data["ActiveMenu"] = "#MenuConnection"
	c.Data["page_title"] = "Details of Connection ID: " + connectionIdStr
	c.TplName = "connection/view.html"
}

func splitUppercase(str string) string {
	regx := regexp.MustCompile(`[A-Z][^A-Z]*`)
	words := regx.FindAllString(str, -1)
	return strings.ToUpper(strings.Join(words, " "))
}
