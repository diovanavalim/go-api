package routes

import (
	"api/src/controller"
	"net/http"
)

var PostRoutes = []Route{
	{
		URI:    "/post",
		Method: http.MethodPost,
		Func:   controller.CreatePost,
		Auth:   true,
	},
	{
		URI:    "/post",
		Method: http.MethodGet,
		Func:   controller.GetPosts,
		Auth:   true,
	},
	{
		URI:    "/post/{id}",
		Method: http.MethodGet,
		Func:   controller.GetPost,
		Auth:   true,
	},
	{
		URI:    "/post/{id}",
		Method: http.MethodPut,
		Func:   controller.UpdatePost,
		Auth:   true,
	},
	{
		URI:    "/post/{id}",
		Method: http.MethodDelete,
		Func:   controller.DeletePost,
		Auth:   true,
	},
}
