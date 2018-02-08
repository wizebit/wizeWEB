package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
	"wizebit/backend/models"
	_ "wizebit/backend/routers"
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
	ormInit()
	beego.Run()
}
