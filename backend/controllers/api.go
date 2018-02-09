package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"strings"
	"encoding/json"
)

type ApiController struct {
	beego.Controller
}

type DeleteRequest struct {
	Filename string
}

//func (a *ApiController) Index() {
//	data := a.Ctx.Input.Data()
//	//beego.Warn(data["exp"], data["publicKey"])
//	publicKey := data["publicKey"].(string)
//
//	var us models.Users
//
//	o := orm.NewOrm()
//	o.Using("default")
//	//find user
//	err := o.QueryTable("users").Filter("public_key", publicKey).Limit(1).One(&us)
//	if err != nil {
//		beego.Error(err)
//
//		a.Ctx.Output.SetStatus(400)
//		a.Data["json"] = map[string]string{"message": err.Error()}
//		a.ServeJSON()
//		a.StopRun()
//		return
//	}
//
//	a.Data["json"] = map[string]string{"hello": "world"}
//	a.ServeJSON()
//	a.StopRun()
//}

func (a *ApiController) GetFileList() {
	data := a.Ctx.Input.Data()
	publicKey := data["publicKey"].(string)

	//	read our storage
	files, err := ioutil.ReadDir("./storage/"+publicKey)
	if err != nil {
		beego.Error(err)

		a.Ctx.Output.SetStatus(500)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}

	//create json with form we need
	var filesSlice []map[string]string
	for _, file := range files {
		delimiter := strings.Index(file.Name(), "~")
		name := file.Name()[delimiter+1:]
		date := file.Name()[:delimiter]

		fileObject := map[string]string{
			"name": name,
			"upload_date": date,
		}
		filesSlice = append(filesSlice, fileObject)
	}

	a.Data["json"] = map[string][]map[string]string{"users_files": filesSlice}
	a.ServeJSON()
	a.StopRun()
}

func (a *ApiController) UploadFile() {
	file, header, err := a.GetFile("file")

	if err != nil {
		beego.Error(err)
		a.Ctx.Output.SetStatus(500)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}

	//	get our user id
	data := a.Ctx.Input.Data()
	publicKey := data["publicKey"].(string)
	//idStr := strconv.Itoa(int(id.(float64)))

	//	read our storage
	folders, err := ioutil.ReadDir("./storage")
	if err != nil {
		beego.Error(err)
		a.Ctx.Output.SetStatus(500)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}

	//	find our user folder
	var fileDirectory string
	for _, folder := range folders {
		if folder.Name() == publicKey {
			fileDirectory = folder.Name()
		}
	}

	//	if there is no folder for user - create it
	if fileDirectory == "" {
		err := os.MkdirAll("./storage/"+publicKey, os.ModePerm)
		if err != nil {
			beego.Error(err)

			a.Ctx.Output.SetStatus(500)
			a.Data["json"] = map[string]string{"message": err.Error()}
			a.ServeJSON()
			a.StopRun()
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
		err := a.SaveToFile("file", "./storage/"+publicKey+"/"+timeNow+"~"+fileName)
		if err != nil {
			beego.Error(err)
		}
	}

	a.Data["json"] = map[string]string{"message": header.Filename+" upload is success"}
	a.ServeJSON()
	a.StopRun()
}

func (a *ApiController) DeleteFile() {
	data := a.Ctx.Input.Data()
	publicKey := data["publicKey"].(string)

	//get body of request
	req := DeleteRequest{}
	json.Unmarshal(a.Ctx.Input.RequestBody, &req)
	//parse body
	if err := a.ParseForm(&req); err != nil {
		beego.Error(err)

		a.Ctx.Output.SetStatus(500)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}

	beego.Warn(req.Filename)

	if _, err := os.Stat("./storage/"+publicKey+"/"+req.Filename); os.IsNotExist(err) {
		beego.Error(err)

		a.Ctx.Output.SetStatus(400)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	} else {
		err := os.Remove("./storage/"+publicKey+"/"+req.Filename)

		if err != nil {
			beego.Error(err)

			a.Ctx.Output.SetStatus(500)
			a.Data["json"] = map[string]string{"message": err.Error()}
			a.ServeJSON()
			a.StopRun()
			return
		}
	}

	a.Data["json"] = map[string]string{"message": req.Filename+" delete is success"}
	a.ServeJSON()
	a.StopRun()
}