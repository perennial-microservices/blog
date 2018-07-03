package service

import . "github.com/perennial-microservices/blog/service"

var AccountRoutes = Routes {
	Route {
		"CreateAccount",
		"POST",
		"/accounts",
		CreateAccount,
	},
	Route {
		"GetAccounts",
		"GET",
		"/accounts",
		GetAccounts,
	},
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		GetAccount,
	},
	Route {
		"UpdateAccount",
		"PUT",
		"/accounts/{accountId}",
		UpdateAccount,
	},
	Route {
		"DeleteAccount",
		"DELETE",
		"/accounts/{accountId}",
		DeleteAccount,
	},
}
