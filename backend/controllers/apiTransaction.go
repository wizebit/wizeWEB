package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"strconv"
)

type TransactionController struct {
	beego.Controller
}

//	transaction json request struct
type Transaction struct {
	From string `json:"from"`
	To string  `json:"to"`
	Amount	string  `json:"amount"`
	MineNow bool  `json:"mineNow"`
}


func (t *TransactionController) responseWithError(status int, message map[string]string, err interface{}) {
	beego.Error(err)

	t.Ctx.Output.SetStatus(status)
	t.Data["json "] = message
	t.ServeJSON()
	t.StopRun()
	return
}

func (t *TransactionController) CreateTransaction() {
	//	get request on transaction
	tr := Transaction{}
	json.Unmarshal(t.Ctx.Input.RequestBody, &tr)

	num, err := strconv.Atoi(tr.Amount)
	if err != nil {
		t.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	//	validation
	if tr.From == "" {
		t.responseWithError(400, map[string]string{"message": "From field is empty"}, "From field is empty!")
		return
	}
	if tr.To == "" {
		t.responseWithError(400, map[string]string{"message": "To field is empty"}, "To field is empty!")
		return
	}
	if num <= 0 {
		t.responseWithError(400, map[string]string{"message": "Amount format is incorrect"}, "Amount format is incorrect: " + tr.Amount)
		return
	}

	//request to blockchain
	req := httplib.Post("http://127.0.0.1:4000/send")

	rawBody := map[string]interface{}{
		"from": tr.From,
		"to": tr.To,
		"amount": num,
		"mineNow": true,
	}

	body, err := json.Marshal(rawBody)
	if err != nil {
		t.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	req.Body(body)
	beego.Warn(req.Body)

	str, err := req.String()
	if err != nil {
		t.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	beego.Warn("response ", str)


	t.Data["json"] = map[string]interface{}{
		"message": "success",
	}
	t.ServeJSON()
	t.StopRun()
}