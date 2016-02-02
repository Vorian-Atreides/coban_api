package main

import (
	"coban/0.1/server/api/controllers/administrations"

	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name		string
	Method		string
	Path		string
	Handler		http.HandlerFunc
}
type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
	return router
}

var routes = Routes {

	//
	//	Common
	//



	//
	//	Administrations
	//

	Route {
		Name: "GetAddresses",
		Method: "GET",
		Path: "/administrations/addresses",
		Handler: administrations.GetAddresses,
	},
	Route {
		Name: "CreateAddress",
		Method: "POST",
		Path: "/administrations/addresses",
		Handler: administrations.CreateAddress,
	},
	Route {
		Name: "GetCompanies",
		Method: "GET",
		Path: "/administrations/companies",
		Handler: administrations.GetCompanies,
	},
	Route {
		Name: "CreateCompany",
		Method: "POST",
		Path: "/administrations/companies",
		Handler: administrations.CreateCompany,
	},
//	Route {
//		Name: "GetAddresses",
//		Method: "GET",
//		Path: "/administrations/addresses",
//		Handler: common.GetAddresses,
//	},
//	Route {
//		Name: "CreateAddress",
//		Method: "POST",
//		Path: "/administrations/addresses",
//		Handler: common.CreateAddress,
//	},
//	Route {
//		Name: "GetCompanies",
//		Method: "GET",
//		Path: "/administrations/companies",
//		Handler: common.GetCompanies,
//	},
//	Route {
//		Name: "CreateCompanies",
//		Method: "POST",
//		Path: "/administrations/companies",
//		Handler: common.CreateCompany,
//	},
}