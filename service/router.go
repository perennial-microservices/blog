package service

import "github.com/gorilla/mux"

func NewRouter(r Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range r {
		router.Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}
	return router
}
