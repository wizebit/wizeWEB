package routers

import (
	"github.com/astaxie/beego"
	"wizebit/backend/controllers"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//	Unauthorised requests
	//
	//	API
	beego.Router("/auth/sign-up", &controllers.AuthController{}, "post:UserSignUp")
	beego.Router("/auth/sign-in", &controllers.AuthController{}, "post:UserSignIn")
	//
	//	Admin panel


	//	Auth requests
	//
	//	API
	//beego.Router("/api/root", &controllers.ApiController{}, "post:Index")
	beego.Router("/api/get-file-list", &controllers.ApiController{}, "get:GetFileList")
	beego.Router("/api/upload-file", &controllers.ApiController{}, "put:UploadFile")
	beego.Router("/api/delete-file", &controllers.ApiController{}, "post:DeleteFile")
	//
	//	Admin panel

	//
	//	Router MiddleWares
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)
}
