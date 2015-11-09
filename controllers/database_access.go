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

  // Not sure if this.

  // User := new(models.User)
  // UserDB := make([]User, resp.Body.numUsers)
  // decoder := json.NewDecoder(resp.Body)
  // decoder.Decode(&UserDB)

  // Update the UserDB.
  UserDB := make([]string, resp.Body.numUsers)
  for i:=0; i<resp.Body.numUsers; i++ {
    UserDB[i] = resp.Body.UserID[i]
  }
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
  // Update the ProductDB.
  ProductDB := make([]string, resp.Body.numItems)
  for i:=0; i<resp.Body.numItems; i++ {
    ProductDB[i] = resp.Body.UserID[i]
  }
  return resp.Body.numItems
}
