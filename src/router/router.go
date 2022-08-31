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
	packageRoutes := routes.UserRoutes

	packageRoutes = append(packageRoutes, routes.LoginRoute)

	for _, route := range packageRoutes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
