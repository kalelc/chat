package main

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

type User struct {
	Name string
	Time string
}

// Session Configuration

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

const SessionName = "LoginSession"

// websocket configuration

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates := template.Must(template.ParseFiles("templates/login.html"))

		if err := templates.ExecuteTemplate(w, "login.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
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
