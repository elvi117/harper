package main

import (
	"harper/auth"
	"net/http"
	"text/template"
)

var tpl *template.Template

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	tpl, _ = template.ParseGlob("./templates/*html")
	// Public handlers
	http.HandleFunc("/logowanie", logInPage)
	http.HandleFunc("/utworz_konto", signUpPage)
	http.HandleFunc("/reset_hasla", resetPasswordPage)
	http.HandleFunc("/utworz_haslo", createPasswordPage)

	// Private handlers
	http.HandleFunc("/", accountStatusPage)
	http.HandleFunc("/konto", myAccountPage)
	http.HandleFunc("/zasady", rulesPage)
	http.HandleFunc("/nagrody", prizesPage)
	http.HandleFunc("/produkty", productsPage)

	// Methods
	http.HandleFunc("/log_in", auth.LogIn)
	http.HandleFunc("/sign_up", auth.SingUp)
	http.HandleFunc("/reset_password", auth.ResetPassword)
	http.HandleFunc("/create_password", auth.CreatePassword)
	http.HandleFunc("/log_out", auth.LogOut)

	http.ListenAndServe(":8080", nil)
}
