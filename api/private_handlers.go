package handler

import "net/http"

func accountStatusPage(w http.ResponseWriter, r *http.Request) {
	checkAuth(w, r, func() {
		tpl.ExecuteTemplate(w, "accountStatus.html", nil)
	})
}

func myAccountPage(w http.ResponseWriter, r *http.Request) {
	checkAuth(w, r, func() {
		tpl.ExecuteTemplate(w, "myAccount.html", nil)
	})
}

func rulesPage(w http.ResponseWriter, r *http.Request) {
	checkAuth(w, r, func() {
		tpl.ExecuteTemplate(w, "rules.html", nil)
	})
}

func prizesPage(w http.ResponseWriter, r *http.Request) {
	checkAuth(w, r, func() {
		tpl.ExecuteTemplate(w, "prizes.html", nil)
	})
}

func productsPage(w http.ResponseWriter, r *http.Request) {
	checkAuth(w, r, func() {
		tpl.ExecuteTemplate(w, "products.html", nil)
	})
}

func checkAuth(w http.ResponseWriter, r *http.Request, handler func()) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/logowanie", http.StatusSeeOther)
		return
	}

	if cookie.Value != "12345" {
		http.Redirect(w, r, "/logowanie", http.StatusSeeOther)
		return
	}
	handler()
}
