package routers

import (
	"github.com/astaxie/beego"
	"wizebit/backend/controllers"
)

func init() {
	//Unauthorised requests
	beego.Router("/auth/sign-up", &controllers.AuthController{}, "post:UserSignUp")
	beego.Router("/auth/sign-in", &controllers.AuthController{}, "post:UserSignIn")
	//Auth requests
	beego.Router("/api/root", &controllers.ApiController{}, "post:Index")

	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
}
