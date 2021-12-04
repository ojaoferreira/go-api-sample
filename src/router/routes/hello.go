package routes

import (
	"api/src/controllers"
	"net/http"
)

var routerHello = []Route{
	{
		URI: "/hello",
		Method: http.MethodGet,
		Func: controllers.Hello,
		Auth: false,
	},
}