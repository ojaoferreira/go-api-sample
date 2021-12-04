package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI string
	Method string
	Func func(http.ResponseWriter, *http.Request)
	Auth bool
}

func Config(r *mux.Router) *mux.Router {
	routes := routerHello

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}