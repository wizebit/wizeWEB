package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"
	"wizeweb/backend/models"
)

type ReportController struct {
	beego.Controller
}

type Report struct {
	Screenshot  string
	Description string
}

func (r *ReportController) responseWithError(status int, message map[string]string, err interface{}) {
	beego.Error(err)

	r.Ctx.Output.SetStatus(status)
	r.Data["json"] = message
	r.ServeJSON()
	r.StopRun()
	return
}

func (r *ReportController) GetReport() {
	//	request info
	report := Report{}
	json.Unmarshal(r.Ctx.Input.RequestBody, &report)

	//	timestamp
	timeNow := strconv.Itoa(int(time.Now().Unix()))

	//	get user id
	data := r.Ctx.Input.Data()
	publicKey := data["publicKey"].(string)
	beego.Warn(publicKey)

	//	get user
	var u models.Users

	o := orm.NewOrm()
	o.Using("default")

	err := o.QueryTable("users").Filter("public_key", publicKey).Limit(1).One(&u)
	if err != nil {
		r.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	//get user id to string
	userId := strconv.Itoa(u.Id)

	//	find our bug-reports folder
	if _, err := os.Stat("./storage/bug-reports/"); os.IsNotExist(err) {
		//	if there is no folder for bug-reports - create it
		err := os.MkdirAll("./storage/bug-reports/", os.ModePerm)
		if err != nil {
			r.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}
	}

	//	create image from base64
	input := report.Screenshot

	b64data := input[strings.IndexByte(input, ',')+1:]

	unbased, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		r.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	br := bytes.NewReader(unbased)
	im, err := png.Decode(br)
	if err != nil {
		r.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	f, err := os.OpenFile("./storage/bug-reports/"+timeNow+"~"+userId+".png", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		r.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	png.Encode(f, im)

	//	write into db
	bugReport := new(models.BugReports)
	bugReport.UserId = u.Id
	bugReport.Message = report.Description
	bugReport.Picture = timeNow + "~" + userId + ".png"

	_, err = o.Insert(bugReport)
	if err != nil {
		r.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	r.Data["json"] = map[string]string{"message": "report is success"}
	r.ServeJSON()
	r.StopRun()
}
