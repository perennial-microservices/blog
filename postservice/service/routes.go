package service

import . "github.com/perennial-microservices/blog/service"

var PostRoutes = Routes {
	Route {
		"CreatePost",
		"POST",
		"/posts",
		CreatePost,
	},
	Route {
		"GetPosts",
		"GET",
		"/posts",
		GetPosts,
	},
	Route{
		"GetPost",
		"GET",
		"/posts/{postId}",
		GetPost,
	},
	Route {
		"UpdatePost",
		"PUT",
		"/posts/{postId}",
		UpdatePost,
	},
	Route {
		"DeletePost",
		"DELETE",
		"/posts/{postId}",
		DeletePost,
	},
}
