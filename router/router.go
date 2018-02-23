package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"../controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{

	//Home page
	Route{
		"Home",
		"GET",
		"/",
		controllers.HomeIndex,
	},
	Route{
		"Transaction PagSeguro",
		"POST",
		"/pagseguro/transaction",
		controllers.TransactionPagSeguro,
	},
	Route{
		"Notification PagSeguro",
		"GET",
		"/pagseguro/notification",
		controllers.NotificationPagSeguro,
	},
	Route{
		"Refirection PagSeguro",
		"GET",
		"/pagseguro/redirection",
		controllers.RedirectionPagSeguro,
	},
}
