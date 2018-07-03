package service

import "net/http"

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

var ApplicationRoutes = Routes{
	Route {
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}


