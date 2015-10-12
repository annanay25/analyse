// Expose API endpoint to add and delete users, items.
package controllers

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
)


func AddUser(w http.ResponseWriter, req *http.Request) {

  fmt.Println("Hello")
  // req.ParseForm()
  log.Println(req)

  example_user:= new(models.User)
  decoder := json.NewDecoder(req.Body)
  decoder.Decode(&example_user)

  fmt.Println("Adding this user to database: ")
  fmt.Println(example_user)
  example_user.AddUserToDatabase()

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
}
