package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"wizeweb/backend/models"
)

type AdminController struct {
	beego.Controller
}

func (a *AdminController) Index() {
	a.TplName = "admin/index.html"

}

func (a *AdminController) ServerList() {
	a.TplName = "admin/serverlist.html"
	o := orm.NewOrm()
	o.Using("default")
	var servers []*models.Servers
	o.QueryTable("servers").OrderBy("id").RelatedSel().Limit(-1).All(&servers)

	a.Data["records"] = servers
}
func (a *AdminController) UsersList() {
	a.TplName = "admin/user.html"
	o := orm.NewOrm()
	o.Using("default")
	var users []*models.Users
	o.QueryTable("users").OrderBy("id").Limit(-1).All(&users)

	a.Data["records"] = users
}
