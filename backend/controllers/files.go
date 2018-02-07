package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
)

type FilesController struct {
	beego.Controller
}

func (f *FilesController) FilesUpload() {
	//data := f.Ctx.Input.Data()

	//beego.Warn(data["exp"], data["customerId"])

	//id := (data["customerId"]).(string)

	files, err := ioutil.ReadDir("./storage")
	if err != nil {
		log.Fatal(err)
	}

	//beego.Warn("id:", id)
	for _, file := range files {
		beego.Warn(file.Name())
	}
}