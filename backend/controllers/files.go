package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"strconv"
	"mime/multipart"
	"time"
)

type Upload struct {
	Data multipart.File
}

type FilesController struct {
	beego.Controller
}

func (f *FilesController) FilesUpload() {
	file, header, err := f.GetFile("file")

	if err != nil {
		beego.Error(err)
		f.Ctx.Output.SetStatus(500)
		f.Data["json"] = map[string]string{"message": err.Error()}
		f.ServeJSON()
		f.StopRun()
		return
	}

	//	get our user id
	data := f.Ctx.Input.Data()
	id := data["customerId"]
	idStr := strconv.Itoa(int(id.(float64)))

	//	read our storage
	folders, err := ioutil.ReadDir("./storage")
	if err != nil {
		beego.Error(err)
		f.Ctx.Output.SetStatus(500)
		f.Data["json"] = map[string]string{"message": err.Error()}
		f.ServeJSON()
		f.StopRun()
		return
	}

	//	find our user folder
	var fileDirectory string
	for _, folder := range folders {
		if folder.Name() == idStr {
			fileDirectory = folder.Name()
		}
	}

	//	if there is no folder for user - create it
	if fileDirectory == "" {
		err := os.MkdirAll("./storage/"+idStr, os.ModePerm)
		if err != nil {
			beego.Error(err)
			f.Ctx.Output.SetStatus(500)
			f.Data["json"] = map[string]string{"message": err.Error()}
			f.ServeJSON()
			f.StopRun()
			return
		}
	}


	if file != nil {
		//	get the filename
		fileName := header.Filename
		//	timestamp
		timeNow := strconv.Itoa(int(time.Now().Unix()))

		//	TODO: Data sharding

		//	save to server
		err := f.SaveToFile("file", "./storage/"+idStr+"/"+timeNow+"~"+fileName)
		if err != nil {
			beego.Error(err)
		}
	}

	f.Data["json"] = map[string]string{"message": header.Filename+" upload is success"}
	f.ServeJSON()
	f.StopRun()
}