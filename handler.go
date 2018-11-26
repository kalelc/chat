package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
	"html/template"
)

type User struct {
	Name string
	Time string
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := User{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/login.html"))

	if name := r.FormValue("name"); name != "" {
		user.Name = name
	}
	if err := templates.ExecuteTemplate(w, "login.html", user); err != nil {
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
