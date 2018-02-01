package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/grrrben/golog"
	"net/http"
	"wizebit/backend/services"
	"wizebit/backend/models"
	"encoding/json"
	"fmt"
)

//info from form will have structure
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//on user sign up (registration)
func (a *App) UserSignUp(w http.ResponseWriter, r *http.Request) {
	raw := services.ReadRequest(w, r)
	var ur User

	// Unmarshal json
	err := json.Unmarshal(raw, &ur)
	if err != nil {
		services.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if ur.Email != "" && ur.Password != "" {
		// bcrypt password
		pas, err := services.HashPassword(ur.Password)

		if err == nil {
			user := models.Users{
				Email: ur.Email,
				Password: pas,
				Role: 20,
			}

			// create user
			o := orm.NewOrm()
			o.Using("default")
			_, err = o.Insert(&user)
			if err != nil {
				golog.Errorf("Unable to create user: %s", err)
				services.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

			services.RespondWithJSON(w, http.StatusOK, map[string]string{"response": "success"})
		} else {
			golog.Errorf("Unable to create user: %s", err)
			services.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	} else if ur.Email == "" {
		services.RespondWithError(w, http.StatusInternalServerError, "Empty email")
	} else if ur.Password == ""{
		services.RespondWithError(w, http.StatusInternalServerError,"Empty password")
	}
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

	if ur.Email != "" && ur.Password != "" {
		user := models.Users{}

		o := orm.NewOrm()
		o.Using("default")
		err := o.QueryTable("users").Filter("email", ur.Email).Limit(1).One(&user)

		if err != nil {
			services.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Println(user)
		match := services.CheckPasswordHash(ur.Password, user.Password)
		fmt.Println("Match:   ", match)

	} else if ur.Email == "" {
		services.RespondWithError(w, http.StatusInternalServerError, "Empty email")
	} else if ur.Password == ""{
		services.RespondWithError(w, http.StatusInternalServerError,"Empty password")
	}
}
