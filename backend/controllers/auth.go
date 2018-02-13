package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/coin-network/curve"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"wizeweb/backend/models"
	"wizeweb/backend/services"
)

type AuthController struct {
	beego.Controller
}

//login json form will have structure
type User struct {
	PrivateKey string `json:"private_key"`
}

func (a *AuthController) responseWithError(status int, message map[string]string, err interface{}) {
	beego.Error(err)

	a.Ctx.Output.SetStatus(status)
	a.Data["json"] = message
	a.ServeJSON()
	a.StopRun()
	return
}

//on user sign up (registration)
func (a *AuthController) UserSignUp() {
	KoblitzCurve := curve.S256() // see https://godoc.org/github.com/btcsuite/btcd/btcec#S256

	privkey, err := curve.NewPrivateKey(KoblitzCurve)

	if err != nil {
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

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
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

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
		a.responseWithError(500, map[string]string{"message": err.Error()}, err)

		return
	}
	//if pk is empty - return error
	if u.PrivateKey == "" {
		a.responseWithError(500, map[string]string{"message": "Empty private key!"}, "Auth: empty private key!")

		return
	}

	us := models.Users{}

	o := orm.NewOrm()
	o.Using("default")
	//find user
	err := o.QueryTable("users").Filter("private_key", services.GetHash(u.PrivateKey)).Limit(1).One(&us)

	if err != nil {
		a.responseWithError(400, map[string]string{"message": err.Error()}, err)

		return
	}

	token, expiresIn, err := services.CreateSignedTokenString(us.PublicKey)
	a.Data["json"] = map[string]interface{}{
		"auth_key":   token,
		"expires_in": expiresIn,
	}
	a.ServeJSON()
	a.StopRun()
}

// customize filters for fine grain authorization
var FilterUser = func(ctx *context.Context) {
	//Unauthorised requests
	if strings.HasPrefix(ctx.Input.URL(), "/auth") || strings.HasPrefix(ctx.Input.URL(), "/storage") {
		return
	}

	//Auth requests
	if strings.HasPrefix(ctx.Input.URL(), "/api") && ctx.Input.Header("Authorization") != "" {
		parsedToken, err := services.ParseTokenFromSignedTokenString(ctx.Input.Header("Authorization"))

		if err == nil && parsedToken.Valid {
			exp := parsedToken.Claims.(jwt.MapClaims)["exp"]
			publicKey := parsedToken.Claims.(jwt.MapClaims)["publicKey"]

			ctx.Input.SetData("exp", exp)
			ctx.Input.SetData("publicKey", publicKey)

			return
		}
	}

	ctx.Output.SetStatus(403)
	ctx.Output.Body([]byte(`{"message": "Unauthorised access to this resource"}`))
}
