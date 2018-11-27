package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

type User struct {
	Name string
	Time string
}

func Login(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/login.html"))

	if err := templates.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/chat", http.StatusSeeOther)
}

func Chat(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/chat.html"))

	user := User{"AN", time.Now().Format(time.Stamp)}
	err := templates.ExecuteTemplate(w, "chat.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
