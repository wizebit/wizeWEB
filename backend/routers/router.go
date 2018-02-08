package routers

import (
	"github.com/astaxie/beego"
	"wizebit/backend/controllers"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//Unauthorised requests
	beego.Router("/auth/sign-up", &controllers.AuthController{}, "post:UserSignUp")
	beego.Router("/auth/sign-in", &controllers.AuthController{}, "post:UserSignIn")
	//Auth requests
	beego.Router("/api/root", &controllers.ApiController{}, "post:Index")
	beego.Router("/api/upload-file", &controllers.FilesController{}, "put:FilesUpload")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
}
