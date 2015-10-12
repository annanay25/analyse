package controllers

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"

  "github.com/mgo"
)

func (user models.User) AddUserToDatabase(){
  session, err = mgo.Dial(127.0.0.1:3000)
  if(err!= nil){
    panic(err)
  }
  // Write to DB.
}
