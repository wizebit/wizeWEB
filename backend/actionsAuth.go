package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/coin-network/curve"
	"github.com/grrrben/golog"
	"net/http"
	"wizebit/backend/models"
	"wizebit/backend/services"
)

type PublicKey ecdsa.PublicKey
type PrivateKey ecdsa.PrivateKey

//login json form will have structure
type User struct {
	PrivateKey string `json:"private_key"`
}

//on user sign up (registration)
func (a *App) UserSignUp(w http.ResponseWriter, r *http.Request) {
	KoblitzCurve := curve.S256() // see https://godoc.org/github.com/btcsuite/btcd/btcec#S256

	privkey, err := curve.NewPrivateKey(KoblitzCurve)

	if err != nil {
		panic("Error on create account")
	}

	pubkey := (privkey.PublicKey)
	address := privkey.PubKey().ToAddress()

	u := models.Users{
		PrivateKey: services.GetHash(privkey.D.String()),
		PublicKey:  services.GetHash(pubkey.X.String()),
		Wallet:     services.GetHash(address),
		Role:       20,
	}

	o := orm.NewOrm()
	o.Using("default")

	_, err = o.Insert(&u)

	if err != nil {
		golog.Errorf("Unable to create user: %s", err)
		services.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	services.RespondWithJSON(
		w,
		http.StatusOK,
		map[string]string{
			"private_key": privkey.D.String(),
			"public_key":  pubkey.X.String(),
			"wallet":      address,
		},
	)
}

//on user sign in (login)
func (a *App) UserSignIn(w http.ResponseWriter, r *http.Request) {
	raw := services.ReadRequest(w, r)
	var ur User

	// Unmarshal json
	err := json.Unmarshal(raw, &ur)
	if err != nil {
		services.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if ur.PrivateKey != "" {
		var u models.Users

		o := orm.NewOrm()
		o.Using("default")
		err := o.QueryTable("users").Filter("private_key", services.GetHash(ur.PrivateKey)).Limit(1).One(&u)

		if err != nil {
			services.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		var role string
		switch u.Role {
		case 0:
			role = "admin"
		case 10:
			role = "moderator"
		default:
			role = "user"
		}
		 //services.CreateTokenMiddleware(w, r, role)

		//services.RespondWithJSON(w, http.StatusOK, u)

		token, err := services.CreateSignedTokenString(role);
		services.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})

	} else if ur.PrivateKey == "" {
		services.RespondWithError(w, http.StatusInternalServerError, "Empty private key")
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	//services.RespondWithError(w, http.StatusInternalServerError, "Gained access to protected resource")

	services.RespondWithJSON(
		w,
		http.StatusOK,
		map[string]string{
			"hello": r.Header.Get("Authorisation"),
		},
	)
}
