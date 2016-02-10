package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"coban/api/src/controllers/administrations"
	"coban/api/src/controllers/clients"
	"coban/api/src/controllers/common"
	"coban/api/src/controllers/offices"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
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

var routes = Routes{

	//
	//	Common
	//

	Route{
		Name:    "Authenticate",
		Method:  "POST",
		Path:    "/users/authenticate",
		Handler: common.AuthenticateRequest,
	},
	Route{
		Name:    "Authenticate",
		Method:  "POST",
		Path:    "/offices/authenticate",
		Handler: common.AuthenticateRequest,
	},
	Route{
		Name:    "Authenticate",
		Method:  "POST",
		Path:    "/administrations/authenticate",
		Handler: common.AuthenticateRequest,
	},

	//
	//	Administrations
	//

	Route{
		Name:    "GetAddresses",
		Method:  "GET",
		Path:    "/administrations/addresses",
		Handler: administrations.GetAddresses,
	},
	Route{
		Name:    "GetCompanies",
		Method:  "GET",
		Path:    "/administrations/companies",
		Handler: administrations.GetCompanies,
	},
	Route{
		Name:    "GetUsers",
		Method:  "GET",
		Path:    "/administrations/users",
		Handler: administrations.GetUsers,
	},
	Route{
		Name:    "GetStations",
		Method:  "GET",
		Path:    "/administrations/stations",
		Handler: administrations.GetStations,
	},
	Route{
		Name:    "GetAccounts",
		Method:  "GET",
		Path:    "/administrations/accounts",
		Handler: administrations.GetAccounts,
	},
	Route{
		Name:    "GetTransportHistories",
		Method:  "GET",
		Path:    "/administrations/transport-histories",
		Handler: administrations.GetTransportHistories,
	},

	//
	// Offices
	//

	Route{
		Name:    "CreateCompany",
		Method:  "POST",
		Path:    "/offices/companies",
		Handler: offices.CreateCompany,
	},
	Route{
		Name:    "GetCurrentCompany",
		Method:  "GET",
		Path:    "/offices/companies",
		Handler: offices.GetCompany,
	},

	Route{
		Name:    "GetEmployees",
		Method:  "GET",
		Path:    "/offices/users",
		Handler: offices.GetEmployees,
	},
	Route{
		Name:    "AddEmployee",
		Method:  "POST",
		Path:    "/offices/users",
		Handler: offices.AddEmployee,
	},
	Route{
		Name:    "UpdateEmployee",
		Method:  "PUT",
		Path:    "/offices/users/{id}",
		Handler: offices.UpdateEmployee,
	},
	Route{
		Name:    "GetTransportHistories",
		Method:  "GET",
		Path:    "/offices/transport-histories",
		Handler: offices.GetTransportHistories,
	},
	Route{
		Name:    "GetTransportHistoriesByUser",
		Method:  "GET",
		Path:    "/offices/transport-histories/{id}",
		Handler: offices.GetTransportHistoryByUser,
	},

	//
	// Clients
	//

	Route{
		Name:    "UpdatePassword",
		Method:  "PUT",
		Path:    "/clients/users",
		Handler: clients.UpdatePassword,
	},
	Route{
		Name:    "GetCurrentUser",
		Method:  "GET",
		Path:    "/clients/users",
		Handler: clients.GetCurrentUser,
	},
	Route{
		Name:    "GetCurrentCompany",
		Method:  "GET",
		Path:    "/clients/companies",
		Handler: clients.GetCurrentCompany,
	},
	Route{
		Name:    "GetTransportHistories",
		Method:  "GET",
		Path:    "/clients/transport-histories",
		Handler: clients.GetTransportHistories,
	},
	Route{
		Name:    "AddTransportHistory",
		Method:  "POST",
		Path:    "/clients/transport-histories",
		Handler: clients.AddTransportHistory,
	},
}
