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
		Auth:   true,
	},
	{
		URI:    "/user/query",
		Method: http.MethodGet,
		Func:   controller.GetUsersByNameOrNickname,
		Auth:   true,
	},
	{
		URI:    "/user/{id}",
		Method: http.MethodGet,
		Func:   controller.GetUserByID,
		Auth:   true,
	},
	{
		URI:    "/user/{id}",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
		Auth:   true,
	},
	{
		URI:    "/user/{id}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
		Auth:   true,
	},
	{
		URI:    "/user/{id}/follow",
		Method: http.MethodPost,
		Func:   controller.FollowUser,
		Auth:   true,
	},
	{
		URI:    "/user/{id}/unfollow",
		Method: http.MethodPost,
		Func:   controller.UnfollowUser,
		Auth:   true,
	},
	{
		URI:    "/user/{id}/followers",
		Method: http.MethodGet,
		Func:   controller.GetUserFollowers,
		Auth:   true,
	},
	{
		URI:    "/user/{id}/following",
		Method: http.MethodGet,
		Func:   controller.GetUserFollowing,
		Auth:   true,
	},
}
