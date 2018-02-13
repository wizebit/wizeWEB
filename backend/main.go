package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
	"wizeweb/backend/models"
	_ "wizeweb/backend/routers"
)

func ormInit() {
	//orm settings
	orm.RegisterDriver("postgres", orm.DRPostgres)
	dbparams := "user=" + beego.AppConfig.String("dbuser") +
		" password=" + beego.AppConfig.String("dbpassword") +
		" host=" + beego.AppConfig.String("dbserver") +
		" port=" + beego.AppConfig.String("dbport") +
		" dbname=" + beego.AppConfig.String("dbname") +
		" sslmode=disable"
	orm.RegisterDataBase("default", "postgres", dbparams)
	orm.Debug = false
	orm.DefaultTimeLoc = time.UTC

	//orm models
	orm.RegisterModel(
		new(models.Users),
	)
}

func main() {
	//public storage
	beego.SetStaticPath("/storage", "storage")
	//orm initiation
	ormInit()
	//beego run action
	beego.Run()
}
