package main

func (a *App) initializeRoutes() {
	//Auth
	a.Router.HandleFunc("/auth/sign-in", a.UserSignIn).Methods("POST")
	a.Router.HandleFunc("/auth/sign-up", a.UserSignUp).Methods("POST")
}
