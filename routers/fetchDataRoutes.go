package routers

import (
	"github.com/annanay25/analyse/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func FetchDataRoutes(router *mux.Router) *mux.Router {
	router.Handle("/addUser",
		negroni.New(
			negroni.HandlerFunc(controllers.AddUser),
		)).Methods("POST")
  router.Handle("/addItem",
		negroni.New(
			negroni.HandlerFunc(controllers.AddItem),
		)).Methods("POST")
  router.Handle("/deleteUser",
		negroni.New(
			negroni.HandlerFunc(controllers.DeleteUser),
		)).Methods("POST")
  router.Handle("/deleteItem",
		negroni.New(
			negroni.HandlerFunc(controllers.DeleteItem),
		)).Methods("POST")

	return router
}
