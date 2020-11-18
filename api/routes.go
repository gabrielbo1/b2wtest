package api

import "net/http"

//Route - Defines structure for mapping the routes that will be mapped by gorilla.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - Array with all mapped APIs.
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Health",
		Method:      "GET",
		Pattern:     "/_health",
		HandlerFunc: Health,
	},
	Route{
		Name:        "Save",
		Method:      "POST",
		Pattern:     "/planet",
		HandlerFunc: Save,
	},
	Route{
		Name:        "Update",
		Method:      "UPDATE",
		Pattern:     "/planet",
		HandlerFunc: Update,
	},
	Route{
		Name:        "GetAll",
		Method:      "GET",
		Pattern:     "/planet",
		HandlerFunc: GetAll,
	},
	Route{
		Name:        "FindByID",
		Method:      "GET",
		Pattern:     "/planet/{id}",
		HandlerFunc: FindByID,
	},
	Route{
		Name:        "FindByName",
		Method:      "GET",
		Pattern:     "/planet/name/{id}",
		HandlerFunc: FindByName,
	},
	Route{
		Name:        "FindByID",
		Method:      "DELETE",
		Pattern:     "/planet/{id}",
		HandlerFunc: FindByID,
	},
}
