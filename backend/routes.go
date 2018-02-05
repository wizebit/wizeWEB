package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"wizebit/backend/services"
)

////PUBLIC ENDPOINTS
//http.HandleFunc("/login", LoginHandler)
//
////PROTECTED ENDPOINTS
//http.Handle("/resource/", negroni.New(
//negroni.HandlerFunc(ValidateTokenMiddleware),
//negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
//))

func (a *App) initializeRoutes() {
	//Unauthorised requests
	a.Router.HandleFunc("/auth/sign-in", a.UserSignIn).Methods("POST")
	a.Router.HandleFunc("/auth/sign-up", a.UserSignUp).Methods("POST")
	//Authorised  requests
	a.Router.Handle("/resource", negroni.New(
		negroni.HandlerFunc(services.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))
}
