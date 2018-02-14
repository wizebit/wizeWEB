package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
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
		new(models.BugReports),
	)
}

func sessionInit() {
	var globalSessions *session.Manager
	sessionConfig := &session.ManagerConfig{
		CookieName: "jigsessionid",
		Gclifetime: 3600,
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func main() {
	//public storage
	beego.SetStaticPath("/storage", "storage")
	//session initiation
	sessionInit()
	//orm initiation
	ormInit()
	//beego run action
	beego.Run()
}
