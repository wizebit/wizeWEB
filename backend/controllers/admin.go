package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (a *AdminController) Index() {
	a.TplName = "admin/index.tpl"
}
