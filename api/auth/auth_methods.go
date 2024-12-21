package auth

import "net/http"

func LogIn(w http.ResponseWriter, r *http.Request) {
	checkPostMethod("/logowanie", w, r, logIn)
}
func SingUp(w http.ResponseWriter, r *http.Request) {
	checkPostMethod("/utworz_konto", w, r, singUp)
}
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	checkPostMethod("/reset_hasla", w, r, resetPassword)
}
func CreatePassword(w http.ResponseWriter, r *http.Request) {
	checkPostMethod("/reset_hasla", w, r, createPassword)
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		cookie.MaxAge = -1
		cookie.Value = ""
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, "/logowanie", http.StatusSeeOther)
}

func logIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "harper@harper.pl" && password == "12345" {
		cookie := http.Cookie{
			Name:     "session",
			Value:    "12345",
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/logowanie", http.StatusSeeOther)
}

func singUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/utworz_konto", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "harper@harper.pl" && password == "12345" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/utworz_konto", http.StatusSeeOther)
}

func resetPassword(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/utworz_haslo", http.StatusSeeOther)
}

func createPassword(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/logowanie", http.StatusSeeOther)
}

func checkPostMethod(redirectUrl string, w http.ResponseWriter, r *http.Request, handler func(http.ResponseWriter, *http.Request)) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
		return
	}
	handler(w, r)
}
