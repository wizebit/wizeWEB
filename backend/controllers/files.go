package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"strconv"
	"os"
)

type FilesController struct {
	beego.Controller
}

func (f *FilesController) FilesUpload() {
	data := f.Ctx.Input.Data()

	id := data["customerId"]

	str:= strconv.Itoa(int(id.(float64)))
	files, err := ioutil.ReadDir("./storage")
	if err != nil {
		beego.Error(err)
	}

	var fileDirectory string

	for _, file := range files {
		if file.Name() == str {
			fileDirectory = file.Name()
		}
	}

	if fileDirectory == "" {
		os.MkdirAll("./storage/"+str, os.ModePerm)
	}


	f.Data["json"] = map[string]string{"message": "success"}
	f.ServeJSON()
	f.StopRun()
}