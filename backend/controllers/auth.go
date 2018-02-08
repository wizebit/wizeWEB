package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/coin-network/curve"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"wizebit/backend/models"
	"wizebit/backend/services"
)

type AuthController struct {
	beego.Controller
}

//login json form will have structure
type User struct {
	PrivateKey string `json:"private_key"`
}

//on user sign up (registration)
func (a *AuthController) UserSignUp() {
	KoblitzCurve := curve.S256() // see https://godoc.org/github.com/btcsuite/btcd/btcec#S256

	privkey, err := curve.NewPrivateKey(KoblitzCurve)

	if err != nil {
		beego.Error(err)

		a.Ctx.Output.SetStatus(500)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}

	pubkey := (privkey.PublicKey)
	address := privkey.PubKey().ToAddress()

	u := new(models.Users)
	u.PrivateKey = services.GetHash(privkey.D.String())
	u.PublicKey = services.GetHash(pubkey.X.String())
	u.Wallet = services.GetHash(address)
	u.Role = 20

	o := orm.NewOrm()
	o.Using("default")

	_, err = o.Insert(u)

	if err != nil {
		beego.Error(err)

		a.Ctx.Output.SetStatus(500)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}

	a.Data["json"] = map[string]string{
		"private_key": privkey.D.String(),
		"public_key":  pubkey.X.String(),
		"wallet":      address,
	}
	a.ServeJSON()
}

//on user sign in (login)
func (a *AuthController) UserSignIn() {
	//get body of request
	u := User{}
	json.Unmarshal(a.Ctx.Input.RequestBody, &u)
	//parse body
	if err := a.ParseForm(&u); err != nil {
		beego.Error(err)

		a.Ctx.Output.SetStatus(500)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}
	//if pk is empty - return error
	if u.PrivateKey == "" {
		a.Ctx.Output.SetStatus(400)
		a.Data["json"] = map[string]string{"message": "Empty private key!"}
		a.ServeJSON()
		a.StopRun()
		return
	}

	us := models.Users{}

	o := orm.NewOrm()
	o.Using("default")
	//find user
	err := o.QueryTable("users").Filter("private_key", services.GetHash(u.PrivateKey)).Limit(1).One(&us)

	if err != nil {
		beego.Error(err)

		a.Ctx.Output.SetStatus(400)
		a.Data["json"] = map[string]string{"message": err.Error()}
		a.ServeJSON()
		a.StopRun()
		return
	}

	token, expiresIn, err := services.CreateSignedTokenString(us.Id)
	a.Data["json"] = map[string]interface{}{
		"auth_key": token,
		"expires_in": expiresIn,
	}
	a.ServeJSON()
	a.StopRun()
}

// customize filters for fine grain authorization
var FilterUser = func(ctx *context.Context) {
	//Unauthorised requests
	if strings.HasPrefix(ctx.Input.URL(), "/auth") {
		return
	}

	//Auth requests
	if strings.HasPrefix(ctx.Input.URL(), "/api") && ctx.Input.Header("Authorization") != "" {
		parsedToken, err := services.ParseTokenFromSignedTokenString(ctx.Input.Header("Authorization"))

		if err == nil && parsedToken.Valid {
			exp := parsedToken.Claims.(jwt.MapClaims)["exp"]
			id := parsedToken.Claims.(jwt.MapClaims)["customerId"]

			//token:=parsedToken.Claims["exp"]
			//fmt.Println(parsedToken.Claims)
			//fmt.Println((parsedToken.Claims.(jwt.MapClaims)["exp"]).(string) + "~" + (parsedToken.Claims.(jwt.MapClaims)["customerRole"]).(string))

			//beego.Info(exp, id)
			ctx.Input.SetData("exp", exp)
			ctx.Input.SetData("customerId", id)

			return
		}
	}

	ctx.Output.SetStatus(403)
	ctx.Output.Body([]byte(`{"message": "Unauthorised access to this resource"}`))
}
