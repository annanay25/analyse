package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = EventListenerRoutes(router)
	router = QueryRoutes(router)
	return router
}
