package routes

import (
	"api/src/controller"
	"net/http"
)

var LoginRoute = Route{
	URI:    "/login",
	Method: http.MethodPost,
	Func:   controller.Login,
	Auth:   false,
}
