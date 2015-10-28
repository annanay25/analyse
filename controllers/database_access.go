package controllers

import (
  "fmt"
  "net/http"
  "log"
  "ioutil"
  "encoding/json"
)

func GetUserInfo() (numUsers int) {
  // Write the database access code here. Basically just send a request to an API that Goutham will write.
  // And ask for all the information through that.

  // Need Number and IDs of users.
  resp, err := http.Get("")   // Add API endpoint.
  if(err!= nil){
    log.Fatal(err)
  }
  defer resp.Body.Close()
  body, error := ioutil.ReadAll(resp.Body)

  return resp.Body.numUsers
}

func GetProductInfo() (numItems int) {
  // Write the database access code here.

  // Need Number and IDs of products.
  resp, err := http.Get("")    // Add API endpoint.
  if(err!= nil){
    log.Fatal(err)
  }
  defer resp.Body.Close()
  body, error := ioutil.ReadAll(resp.Body)

  return resp.Body.numItems
}
