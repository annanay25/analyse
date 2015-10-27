package controllers

import (
  "fmt"
  "net/http"
  "log"
  "ioutil"
  "encoding/json"
)

func GetUserInfo(){
  // Write the database access code here. Basically just send a request to an API that Goutham will write.
  // And ask for all the information through that.

  // Just need number of users.
  resp, err := http.Get("")   // Add API endpoint.
  if(err!= nil){
    log.Fatal(err)
  }
  defer resp.Body.Close()
  body, error := ioutil.ReadAll(resp.Body)

  return resp.Body.numUsers
}

func GetProductInfo(){
  // Write the database access code here.

  // Just need number of items.
  resp, err := http.Get("")    // Add API endpoint.
  if(err!= nil){
    log.Fatal(err)
  }
  defer resp.Body.Close()
  body, error := ioutil.ReadAll(resp.Body)

  return resp.Body.numItems
}
