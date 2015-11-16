package controllers

import (
  "fmt"
  "net"
  "net/http"
)

func QueryController(w http.ResponseWriter, req *http.Request){
  // Query handling comes here.
  query:= new(models.Query)
  decoder := json.NewDecoder(req.Body)
  decoder.Decode(&query)
  // Calculate the xTy matrix products and get the top recommendations.

}


func TopNRec(){

}
