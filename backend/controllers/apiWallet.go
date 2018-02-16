package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
)

type WalletController struct {
	beego.Controller
}

type WalletList struct {
	ListWallets []string `json:"listWallets"`
	Success bool `json:"success"`
}

type Wallet struct {
	Credit int  `json:"credit"`
	Success bool `json:"success"`
}

func (w *WalletController) responseWithError(status int, message map[string]string, err interface{}) {
	beego.Error(err)

	w.Ctx.Output.SetStatus(status)
	w.Data["json"] = message
	w.ServeJSON()
	w.StopRun()
	return
}

func (w *WalletController) WalletsList() {
	req := httplib.Get("http://127.0.0.1:4000/wallets/list")

	str, err := req.String()
	if err != nil {
		w.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	var response WalletList

	err = json.Unmarshal([]byte(str), &response)
	if err != nil {
		w.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	w.Data["json"] = map[string]interface{}{
		"walletsList": response.ListWallets,
	}
	w.ServeJSON()
	w.StopRun()
}

func (w *WalletController) WalletCheck() {
	walletNumber := w.GetString(":walletNumber")
	req := httplib.Get("http://127.0.0.1:4000/wallet/"+walletNumber)

	str, err := req.String()
	if err != nil {
		w.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	var response Wallet

	err = json.Unmarshal([]byte(str), &response)
	if err != nil {
		w.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	w.Data["json"] = map[string]interface{}{
		"walletInfo": map[string]interface{}{
			"credit":  response.Credit,
			"success": response.Success,
		},
	}
	w.ServeJSON()
	w.StopRun()
}