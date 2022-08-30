package routes

import (
	"api/src/controller"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:    "/user",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
		Auth:   false,
	},
	{
		URI:    "/user",
		Method: http.MethodGet,
		Func:   controller.GetUser,
		Auth:   false,
	},
	{
		URI:    "/user/query",
		Method: http.MethodGet,
		Func:   controller.GetUsersByNameOrNickname,
		Auth:   false,
	},
	{
		URI:    "/user/{id}",
		Method: http.MethodGet,
		Func:   controller.GetUserByID,
		Auth:   false,
	},
	{
		URI:    "/user/{id}",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
		Auth:   false,
	},
	{
		URI:    "/user/{id}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
		Auth:   false,
	},
}
