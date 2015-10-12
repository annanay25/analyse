package main

import (
  "github.com/annanay25/analyse/routers"
  "github.com/codegangsta/negroni"
  "net/http"
)

func main() {
  router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}
