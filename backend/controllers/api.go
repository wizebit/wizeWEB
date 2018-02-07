package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (a *ApiController) Index() {
	data := a.Ctx.Input.Data()
	beego.Warn(data["exp"], data["customerId"])

	a.Data["json"] = map[string]string{"hello": "world"}
	a.ServeJSON()
	a.StopRun()
}
