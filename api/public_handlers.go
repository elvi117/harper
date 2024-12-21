package handler

import "net/http"

func logInPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func signUpPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "signUp.html", nil)
}

func resetPasswordPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "resetPassword.html", nil)
}

func createPasswordPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "createPassword.html", nil)
}
