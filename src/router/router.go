package router

import (
	"api/src/middleware"
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
	packageRoutes = append(packageRoutes, routes.PostRoutes...)

	for _, route := range packageRoutes {
		if route.Auth {
			r.HandleFunc(route.URI, middleware.Logger(middleware.Authenticate(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
