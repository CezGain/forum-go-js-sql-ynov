package forum

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./static/home/index.html")
		t.Execute(w, "hello world")
	}
}

// tmpl, _ := template.ParseFiles("./index.html")
// 	tmpl.Execute(w, nil)

func Connexion_Creation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !ValidSession(w, r) {
			t, _ := template.ParseFiles("./static/connexion-creation/index.html")
			t.Execute(w, "hello world")
		}
		//  else {
		// w.Header().Set("content-type", "text/html; charset=utf-8")
		// http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		// }
	}
}

func Moderation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./static/admin/index.html")
		t.Execute(w, "hello world")
	}
}
func Profil() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./static/profil/index.html")
		t.Execute(w, "hello world")
		username := mux.Vars(r)["nameUser"]
		fmt.Println(username)
	}
}
