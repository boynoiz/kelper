package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	//Route{
	//	"Auth",
	//	"POST",
	//	"/auth",
	//	JWTAuth,
	//},
	Route{
		"Ping",
		"GET",
		"/ping",
		PingHandler,
	},
	Route{
		"Shout",
		"POST",
		"/shout",
		ShoutHandler,
	},
	//Route{
	//	"Execute",
	//	"POST",
	//	"/exe/{projectName}/{branchName}",
	//	ExecuteHandler,
	//},
}
