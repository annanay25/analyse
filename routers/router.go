package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = fetchDataRoutes(router)
	router = queryRoutes(router)
	return router
}
