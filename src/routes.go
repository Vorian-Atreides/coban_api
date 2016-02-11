package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"coban/api/src/controllers/administrations"
	"coban/api/src/controllers/clients"
	"coban/api/src/controllers/common"
	"coban/api/src/controllers/offices"
)

type route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}
type routes []route

// NewRouter generate the routes for the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range items {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
	return router
}

var items = routes{

	//
	//	Common
	//

	route{
		Name:    "Authenticate",
		Method:  "POST",
		Path:    "/users/authenticate",
		Handler: common.AuthenticateRequest,
	},
	route{
		Name:    "Authenticate",
		Method:  "POST",
		Path:    "/offices/authenticate",
		Handler: common.AuthenticateRequest,
	},
	route{
		Name:    "Authenticate",
		Method:  "POST",
		Path:    "/administrations/authenticate",
		Handler: common.AuthenticateRequest,
	},

	//
	//	Administrations
	//

	route{
		Name:    "GetAddresses",
		Method:  "GET",
		Path:    "/administrations/addresses",
		Handler: administrations.GetAddresses,
	},

	route{
		Name:    "GetCompanies",
		Method:  "GET",
		Path:    "/administrations/companies",
		Handler: administrations.GetCompanies,
	},

	route{
		Name:    "GetUsers",
		Method:  "GET",
		Path:    "/administrations/users",
		Handler: administrations.GetUsers,
	},

	route{
		Name:    "GetStations",
		Method:  "GET",
		Path:    "/administrations/stations",
		Handler: administrations.GetStations,
	},

	route{
		Name:    "GetAccounts",
		Method:  "GET",
		Path:    "/administrations/accounts",
		Handler: administrations.GetAccounts,
	},

	route{
		Name:    "GetTransportHistories",
		Method:  "GET",
		Path:    "/administrations/transport-histories",
		Handler: administrations.GetTransportHistories,
	},

	//
	// Offices
	//

	route{
		Name:    "CreateCompany",
		Method:  "POST",
		Path:    "/offices/companies",
		Handler: offices.CreateCompany,
	},
	route{
		Name:    "GetCurrentCompany",
		Method:  "GET",
		Path:    "/offices/companies",
		Handler: offices.GetCompany,
	},

	route{
		Name:    "GetEmployees",
		Method:  "GET",
		Path:    "/offices/users",
		Handler: offices.GetEmployees,
	},
	route{
		Name:    "AddEmployee",
		Method:  "POST",
		Path:    "/offices/users",
		Handler: offices.AddEmployee,
	},
	route{
		Name:    "UpdateEmployee",
		Method:  "PUT",
		Path:    "/offices/users/{id}",
		Handler: offices.UpdateEmployee,
	},

	route{
		Name:    "GetTransportHistories",
		Method:  "GET",
		Path:    "/offices/transport-histories",
		Handler: offices.GetTransportHistories,
	},
	route{
		Name:    "GetTransportHistoriesByUser",
		Method:  "GET",
		Path:    "/offices/transport-histories/{id}",
		Handler: offices.GetTransportHistoryByUser,
	},

	route{
		Name:    "GetAddresses",
		Method:  "GET",
		Path:    "/offices/addresses",
		Handler: offices.GetAddresses,
	},
	route{
		Name:    "AddAddress",
		Method:  "POST",
		Path:    "/offices/addresses",
		Handler: offices.AddAddress,
	},

	//
	// Clients
	//

	route{
		Name:    "GetCurrentUser",
		Method:  "GET",
		Path:    "/clients/users",
		Handler: clients.GetCurrentUser,
	},
	route{
		Name:    "UpdatePassword",
		Method:  "PUT",
		Path:    "/clients/users",
		Handler: clients.UpdatePassword,
	},

	route{
		Name:    "GetCurrentCompany",
		Method:  "GET",
		Path:    "/clients/companies",
		Handler: clients.GetCurrentCompany,
	},

	route{
		Name:    "GetTransportHistories",
		Method:  "GET",
		Path:    "/clients/transport-histories",
		Handler: clients.GetTransportHistories,
	},
	route{
		Name:    "AddTransportHistory",
		Method:  "POST",
		Path:    "/clients/transport-histories",
		Handler: clients.AddTransportHistory,
	},
}
