package routers

import (
	"github.com/annanay25/analyse/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func AddUserRoutes(router *mux.Router) *mux.Router {
	router.Handle("/AddUser",
		negroni.New(
			negroni.HandlerFunc(controllers.AddUser),
		)).Methods("POST")

	return router
}

func AddItemRoutes(router *mux.Router) *mux.Router {
	router.Handle("/AddItem",
		negroni.New(
			negroni.HandlerFunc(controllers.AddItem),
		)).Methods("POST")

	return router
}

func EventListenerRoutes(router *mux.Router) *mux.Router {
	router.Handle("/event",
		negroni.New(
			negroni.HandlerFunc(controllers.EventController),
		)).Methods("POST")

	return router
}

func QueryRoutes(router *mux.Router) *mux.Router {
	router.Handle("/query",
		negroni.New(
			negroni.HandlerFunc(controllers.QueryController),
		)).Methods("GET")

	return router
}
