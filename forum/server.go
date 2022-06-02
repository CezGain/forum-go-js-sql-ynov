package forum

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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
		t, _ := template.ParseFiles("./static/connexion-creation/index.html")
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

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u UserLogin
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &u)
		fmt.Println(string(body))
		fmt.Println(u)
		// http.Redirect(w, r, "/home", http.StatusFound)
	}
}
