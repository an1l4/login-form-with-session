package accountcontroller

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
	tmp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if username == "anila" && password == "123" {
		session, _ := store.Get(r, "mysession")
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/account/welcome", http.StatusSeeOther)

	} else {
		data := map[string]interface{}{
			"err": "Invalid",
		}
		tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
		tmp.Execute(w, data)

	}

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	tmp, _ := template.ParseFiles("views/accountcontroller/welcome.html")
	tmp.Execute(w, data)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/account/index", http.StatusSeeOther)
}
