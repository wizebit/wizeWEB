package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"wizeweb/backend/models"
	"wizeweb/backend/services"
	"encoding/binary"
	"github.com/astaxie/beego/httplib"
)

type AuthController struct {
	beego.Controller
}

var sessionName = beego.AppConfig.String("SessionName")

//	register form json structure
type UserSignUp struct {
	PrivateKey string `json:"privkey"`
	PublicKey string `json:"pubkey"`
	Address string `json:"address"`
	AesKey string `json:"aesKey"`
}

//	login json form will have structure
type User struct {
	PublicKey string `json:"publicKey"`
	AesKey string `json:"aesKey"`
}

//	auth form for admin panel
type Admin struct {
	PublicKey string `form:"public_key"`
	AesKey string `form:"aes_key"`
}

func (a *AuthController) responseWithError(status int, message map[string]string, err interface{}) {
	beego.Error(err)

	a.Ctx.Output.SetStatus(status)
	a.Data["json"] = message
	a.ServeJSON()
	a.StopRun()

	return
}

//func (a *AuthController) UserPreSignUp(){
//	var w *wallet.Wallet
//
//	w = wallet.NewWallet()
//
//	a.Data["json"] = map[string]interface{}{
//		"privateKey": w.PrivateKey.D.String(),
//		"publicKey": w.PublicKey,
//		"address": w.GetAddress(),
//	}
//	a.ServeJSON()
//	a.StopRun()
//}
//
//func (a *AuthController) UserSignUp() {
//	//	get body of request
//	uf := UserSignUp{}
//	json.Unmarshal(a.Ctx.Input.RequestBody, &uf)
//
//	if binary.Size([]byte(uf.AesKey)) != 32 {
//		a.responseWithError(400, map[string]string{"message": "password is not equal to 32 bytes"}, "password is not equal to 32 bytes. AesKey length:" + string(binary.Size([]byte(uf.AesKey))))
//
//		return
//	}
//
//	//	Encode private key
//	csk, err := services.GetAESEncode(uf.PrivateKey, uf.AesKey)
//	if err != nil {
//		a.responseWithError(500, map[string]string{"message": err.Error()}, err)
//
//		return
//	}
//
//	//	create user in DB
//	u := new(models.Users)
//	u.PrivateKey = csk
//	u.PublicKey = uf.PublicKey
//	u.Address = uf.Address
//	u.Role = 20
//
//	o := orm.NewOrm()
//	o.Using("default")
//
//	_, err = o.Insert(u)
//
//	if err != nil {
//		a.responseWithError(500, map[string]string{"message": err.Error()}, err)
//
//		return
//	}
//
//	//	return result
//	a.Data["json"] = map[string]interface{}{
//		"message": "success",
//	}
//	a.ServeJSON()
//	a.StopRun()
//}

func (a *AuthController) SignUp() {
		//	get body of request (for aes key)
		uf := UserSignUp{}
		json.Unmarshal(a.Ctx.Input.RequestBody, &uf)

		if uf.AesKey == "" {
			a.responseWithError(400, map[string]string{"message": "password is empty"}, "password is empty!")

			return
		} else if binary.Size([]byte(uf.AesKey)) != 32 {
			a.responseWithError(400, map[string]string{"message": "password is not equal to 32 bytes"}, "password is not equal to 32 bytes. AesKey length:" + string(binary.Size([]byte(uf.AesKey))))

			return
		}

		//	register wallet in blockchain
		req := httplib.Post("http://127.0.0.1:4000/wallet/new")

		str, err := req.String()
		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}

		//	get credentials
		ub := UserSignUp{}
		err = json.Unmarshal([]byte(str), &ub)
		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}

		//	Encode private key
		csk, err := services.GetAESEncode(ub.PrivateKey, uf.AesKey)
		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}

		//	create user in DB
		u := new(models.Users)
		u.PrivateKey = csk
		u.PublicKey = ub.PublicKey
		u.Address = ub.Address
		u.Role = 20

		o := orm.NewOrm()
		o.Using("default")

		_, err = o.Insert(u)

		if err != nil {
			a.responseWithError(500, map[string]string{"message": err.Error()}, err)

			return
		}

		//	return result
		a.Data["json"] = map[string]interface{}{
			"privateKey":	ub.PrivateKey,
			"publicKey": ub.PublicKey,
			"address": ub.Address,
			"password": uf.AesKey,
		}
		a.ServeJSON()
		a.StopRun()
}

func (a *AuthController) UserSignIn() {
	//	get body of request
	uf := User{}
	json.Unmarshal(a.Ctx.Input.RequestBody, &uf)

	//	if pk is empty - return error
	if uf.PublicKey == "" {
		a.responseWithError(500, map[string]string{"message": "Empty private key!"}, "Auth: empty private key!")

		return
	} else if uf.AesKey == "" {
		a.responseWithError(500, map[string]string{"message": "Empty password!"}, "Auth: empty password!")

		return
	}

	u := models.Users{}

	o := orm.NewOrm()
	o.Using("default")

	//	find user
	err := o.QueryTable("users").Filter("public_key", uf.PublicKey).Limit(1).One(&u)

	if err != nil {
		a.responseWithError(400, map[string]string{"message": err.Error()}, err)

		return
	}

	//	decode Private Key
	csk, err := services.GetAESDecode(u.PrivateKey, uf.AesKey)

	if err != nil {
		a.responseWithError(400, map[string]string{"message": err.Error()}, err)

		return
	}

	hashKey := services.GetHash(csk)
	u.SessionKey = hashKey
	_, err = o.Update(&u)
	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}

	token, expiresIn, err := services.CreateSignedTokenString(hashKey)
	a.Data["json"] = map[string]interface{}{
		"accessToken":   token,
		"expiresIn": expiresIn,
	}
	a.ServeJSON()
	a.StopRun()
}

//	admin sign in
func (a *AuthController) AdminAuth() {
	a.TplName = "auth/index.tpl"
}

func (a *AuthController) AdminSignIn() {
	adm := Admin{}
	if err := a.ParseForm(&adm); err != nil {
		beego.Error(err)
		a.Data["errorMessage"] = err.Error()
	}

	if adm.PublicKey == "" || adm.AesKey == "" || binary.Size([]byte(adm.AesKey)) != 32 {
		if adm.PublicKey == "" {
			beego.Error("Your public key is empty")
			a.Data["errorMessage"] = "Your public key is empty"
		}
		if adm.AesKey == "" {
			beego.Error("Your aes key is empty")
			a.Data["errorMessage2"] = "Your password is empty"
		} else if binary.Size([]byte(adm.AesKey)) != 32 {
			a.responseWithError(400, map[string]string{"message": "password is not equal to 32 bytes"}, "password is not equal to 32 bytes. AesKey length:"+string(binary.Size([]byte(adm.AesKey))))
		}

		a.TplName = "auth/index.tpl"
		return
	}

	var u models.Users

	o := orm.NewOrm()
	o.Using("default")

	//	find user
	err := o.QueryTable("users").Filter("public_key", adm.PublicKey).Limit(1).One(&u)

	if err != nil {
		beego.Error(err)
		a.Data["errorMessage"] = err.Error()
		a.TplName = "auth/index.tpl"
		return
	}

	//	decode Private Key
	csk, err := services.GetAESDecode(u.PrivateKey, adm.AesKey)

	if err != nil {
		beego.Error(err)
		a.Data["errorMessage2"] = err.Error()
		a.TplName = "auth/index.tpl"
		return
	}

	hashKey := services.GetHash(csk)
	u.SessionKey = hashKey
	_, err = o.Update(&u)
	if err != nil {
		beego.Error(err)
		a.TplName = "auth/index.tpl"
		return
	}

	if u.Role == 0 {
		v := a.GetSession(sessionName)
		if v == nil {
			a.SetSession(sessionName, u.PublicKey)
		}
		a.Redirect("/admin", 302)
	}

	a.Data["errorMessage"] = "Unauthorised access to this resource"
	a.TplName = "auth/index.tpl"
}
func (a *AuthController) AdminSignOut() {
	a.DelSession(sessionName)
	a.Redirect("/auth/admin", 302)
}

//	customize filters for fine grain authorization
var FilterUser = func(ctx *context.Context) {
	//	Unauthorised requests
	if strings.HasPrefix(ctx.Input.URL(), "/auth") || strings.HasPrefix(ctx.Input.URL(), "/storage") {
		return
	}

	//	Auth requests
	if strings.HasPrefix(ctx.Input.URL(), "/api") && ctx.Input.Header("X-ACCESS-TOKEN") != "" {
		parsedToken, err := services.ParseTokenFromSignedTokenString(ctx.Input.Header("X-ACCESS-TOKEN"))

		if err == nil && parsedToken.Valid {
			expiresIn := parsedToken.Claims.(jwt.MapClaims)["expiresIn"]
			hashKey := parsedToken.Claims.(jwt.MapClaims)["hashKey"]

			ctx.Input.SetData("expiresIn", expiresIn)
			ctx.Input.SetData("hashKey", hashKey)

			return
		}
	}

	if strings.HasPrefix(ctx.Input.URL(), "/admin") {
		_, ok := ctx.Input.Session(sessionName).(string)

		if ok {
			var u models.Users
			o := orm.NewOrm()
			o.Using("default")

			//	find user
			err := o.QueryTable("users").Filter("public_key", ctx.Input.Session(sessionName).(string)).Limit(1).
				One(&u)

			if err == nil && u.Role == 0 {
				return
			} else {
				ctx.Redirect(302, "/auth/admin")
			}
		} else {
			ctx.Redirect(302, "/auth/admin")
		}
	}

	ctx.Output.SetStatus(403)
	ctx.Output.Body([]byte(`{"message": "Unauthorised access to this resource"}`))
}