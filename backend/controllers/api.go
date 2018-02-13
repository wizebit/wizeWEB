package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"wizeweb/backend/services"
)

type ApiController struct {
	beego.Controller
}

type FileRequest struct {
	Filename   string
	TransferTo string
}

func (a *ApiController) responseWithError(status int, message map[string]string, err interface{}) {
	beego.Error(err)

	a.Ctx.Output.SetStatus(status)
	a.Data["json"] = message
	a.ServeJSON()
	a.StopRun()
	return
}

func (a *ApiController) GetFileList() {
	data := a.Ctx.Input.Data()
	publicKey := services.Trim(data["publicKey"].(string))

	//	find our user folder
	if _, err := os.Stat("./storage/" + publicKey); os.IsNotExist(err) {
		//	if there is no folder for user - create it
		err := os.MkdirAll("./storage/"+publicKey, os.ModePerm)
		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}
	}

	//	read our storage
	files, err := ioutil.ReadDir("./storage/" + publicKey)
	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	//	create json with form we need
	var filesSlice []map[string]string
	for _, file := range files {
		delimiter := strings.Index(file.Name(), "~")
		name := file.Name()[delimiter+1:]
		date := file.Name()[:delimiter]
		rel := "/storage/" + publicKey + "/" + file.Name()
		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}

		fileObject := map[string]string{
			"name":        name,
			"uploadDate": date,
			"relativePath": rel,
		}
		filesSlice = append(filesSlice, fileObject)
	}

	a.Data["json"] = map[string][]map[string]string{"userFiles": filesSlice}
	a.ServeJSON()
	a.StopRun()
}

func (a *ApiController) UploadFile() {
	file, header, err := a.GetFile("file")

	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	//	get our user id
	data := a.Ctx.Input.Data()
	publicKey := services.Trim(data["publicKey"].(string))
	//idStr := strconv.Itoa(int(id.(float64)))

	//	find our user folder
	if _, err := os.Stat("./storage/" + publicKey); os.IsNotExist(err) {
		//	if there is no folder for user - create it
		err := os.MkdirAll("./storage/"+publicKey, os.ModePerm)
		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

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

	a.Data["json"] = map[string]string{"message": header.Filename + " upload is success"}
	a.ServeJSON()
	a.StopRun()
}

func (a *ApiController) DeleteFile() {
	data := a.Ctx.Input.Data()
	publicKey := services.Trim(data["publicKey"].(string))

	//get body of request
	req := FileRequest{}
	json.Unmarshal(a.Ctx.Input.RequestBody, &req)

	//parse body
	if err := a.ParseForm(&req); err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	if _, err := os.Stat("./storage/" + publicKey + "/" + req.Filename); os.IsNotExist(err) {
		a.responseWithError(400, map[string]string{"message": err.Error()}, err)

		return
	}

	err := os.Remove("./storage/" + publicKey + "/" + req.Filename)

	if err != nil {
		a.responseWithError(400, map[string]string{"message": err.Error()}, err)

		return
	}

	a.Data["json"] = map[string]string{"message": req.Filename + " delete is success"}
	a.ServeJSON()
	a.StopRun()
}

func (a *ApiController) TransferFile() {
	data := a.Ctx.Input.Data()
	publicKey := services.Trim(data["publicKey"].(string))

	//	get body of request
	req := FileRequest{}
	json.Unmarshal(a.Ctx.Input.RequestBody, &req)

	//	parse body
	if err := a.ParseForm(&req); err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	beego.Warn(req)

	//	copy file with check if file/folders exist
	if _, err := os.Stat("./storage/" + publicKey + "/" + req.Filename); os.IsNotExist(err) {
		a.responseWithError(400, map[string]string{"message": err.Error()}, err)

		return
	}

	//	check if user who will get transfer is exist
	o := orm.NewOrm()
	o.Using("default")
	exist := o.QueryTable("users").Filter("publicKey", services.GetHash(req.TransferTo)).Exist()
	if !exist {
		a.responseWithError(400, map[string]string{"message": "there is no such user"},
		"Transfer file: there is no such user")

		return
	}

	//	find our user folder
	if _, err := os.Stat("./storage/" + services.GetHash(req.TransferTo)); os.IsNotExist(err) {
		//	if there is no folder for user - create it
		err := os.MkdirAll("./storage/"+services.GetHash(req.TransferTo), os.ModePerm)
		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}
	}

	//	start copy stream
	//
	//	from
	from, err := os.Open("./storage/" + publicKey + "/" + req.Filename)
	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}
	defer from.Close()
	//	to
	to, err := os.OpenFile("./storage/"+services.GetHash(req.TransferTo)+"/"+req.Filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}
	defer to.Close()
	//	copy
	_, err = io.Copy(to, from)
	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	//	remove old file
	err = os.Remove("./storage/" + publicKey + "/" + req.Filename)

	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	a.Data["json"] = map[string]string{"message": req.Filename + " transfer to " + req.TransferTo + " is success"}
	a.ServeJSON()
	a.StopRun()
}
