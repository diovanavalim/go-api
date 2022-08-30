package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	return ConfigRouter(router)
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routes := routes.UserRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
