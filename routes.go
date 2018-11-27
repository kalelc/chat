package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	return router
}

var routes = Routes{
	Route{
		"login",
		"/",
		"GET",
		Login,
	},
	Route{
		"login",
		"/login",
		"POST",
		Login,
	},
	Route{
		"chat",
		"/chat",
		"GET",
		Chat,
	},
}
