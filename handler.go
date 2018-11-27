package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
	"os"
	"github.com/gorilla/sessions"
)

type User struct {
	Name string
	Time string
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const SessionName = "LoginSession"

func Login(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/login.html"))

	if err := templates.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, SessionName)

	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	session.Values["name"] = r.FormValue("name")
	session.Values["datetime"] = time.Now().Format(time.Stamp)
	session.Save(r, w)

	http.Redirect(w, r, "/chat", http.StatusSeeOther)
}

func Chat(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/chat.html"))

	session, err := store.Get(r, SessionName)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


	if name, ok := session.Values["name"].(string); ok {
		datetime := session.Values["datetime"].(string)
		user := User{name, datetime}
		if err := templates.ExecuteTemplate(w, "chat.html", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func logRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return
	}

	log.Printf("%s", dump)
}

func jsonBody(req *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(req.Body)

	mapper := make(map[string]interface{})
	err := decoder.Decode(&mapper)

	if err != nil {
		log.Println(err)
	}
	return mapper, err
}
