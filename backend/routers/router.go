package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"wizeweb/backend/controllers"
)

func init() {
	//	Unauthorised requests
	//
	//	API
	beego.Router("/auth/sign-up", &controllers.AuthController{}, "post:UserSignUp")
	beego.Router("/auth/sign-in", &controllers.AuthController{}, "post:UserSignIn")
	beego.Router("/auth/admin", &controllers.AuthController{}, "get:AdminAuth")
	beego.Router("/auth/admin/sign-in", &controllers.AuthController{}, "post:AdminSignIn")
	//
	//	Admin panel

	//	Auth requests
	//
	//	API
	//	Files
	beego.Router("/api/get-files-list", &controllers.ApiController{}, "get:GetFilesList")
	beego.Router("/api/upload-file", &controllers.ApiController{}, "put:UploadFile")
	beego.Router("/api/delete-file", &controllers.ApiController{}, "post:DeleteFile")
	beego.Router("/api/transfer-file", &controllers.ApiController{}, "post:TransferFile")
	// Bug report
	beego.Router("/api/report-bug", &controllers.ReportController{}, "post:GetReport")
	//	Wallets
	beego.Router("/api/get-wallets-list", &controllers.WalletController{}, "get:WalletsList")
	beego.Router("/api/wallet/:walletNumber", &controllers.WalletController{}, "get:WalletCheck")

	//	Admin panel
	beego.Router("/admin", &controllers.AdminController{}, "get:Index")
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
