// Expose API endpoint to add and delete users, items.
package controllers

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"

  "github.com/annanay25/analyse/models"
  "github.com/skelterjohn/go.matrix"
)


func EventController(w http.ResponseWriter, req *http.Request) {

  // Implementing the events server. Code for listening to the events will come here.
  // Events are going to be view / buy events, with 'weightToAdd'. Add this to the Ratings matrix.

  fmt.Println("Hello")
  log.Println(req)

  example_event:= new(models.Event)
  decoder := json.NewDecoder(req.Body)
  decoder.Decode(&example_event)

  fmt.Println("Event: ")
  fmt.Println(example_event)
  // Update values in ALS.Ratings_matrix
  userID := example_event.userID
  productID := example_event.productID
  weightToAdd := example_event.weightToAdd

  // Get which user this userID belongs to.
  i := 0
  for controllers.UserDB[i] != userID {
    i ++
  }
  userNumber := i

  i = 0
  for controllers.ProductDB[i] != productID {
    i++
  }
  productNumber := i
  // Now add this to the Ratings matrix.
  ALS.Ratings_Matrix.Set(userNumber, productNumber, weightToAdd)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOk)
}

func AddUser(w http.ResponseWriter, req *http.Request){
  // Increase number of columns in the User Matrix 'X'
}

func AddItem(w http.ResponseWriter, req *http.Request){
  // Increase number of columns in the Item Matrix 'Y'
}
