// API endpoint for event queue consumption.
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

  log.Println(req)

  example_event:= new(models.Event)
  decoder := json.NewDecoder(req.Body)
  decoder.Decode(&example_event)

  fmt.Println("Event: ")
  fmt.Println(example_event)
  // Update values in ALS.Ratings_matrix
  UserID := example_event.UserID
  ProductID := example_event.ProductID
  Event := example_event.Event

  // Calculate weight to add.
  if Event == "View Product" {
    weightToAdd = 0.1
  }
  if Event == "Add Product to Cart" {
    weightToAdd = 0.5
  }
  if Event == "Delete from Cart" {
    weightToAdd = 0.2
  }

  // Get which user this userID belongs to.
  i := 0
  if NumUsers==0 {
    UserDB := make([]string, 1)
    UserDB[0] = UserID
    NumUsers = 1
    ALS.AddUser()
  }
  for i=0; i<NumUsers; i++ {
    if UserDB[i] = UserID {
      break
    }
  }
  if(i == NumUsers){
    // User does not exist. Add to UserDB.
    UserDB = append(UserDB, UserID)
    NumUsers = NumUsers+1
    ALS.AddUser()
  }
  userNumber := i

  i = 0
  if NumItems==0 {
    ProductDB := make([]string, 1)
    ProductDB[0] = ProductID
    NumItems = 1
    ALS.AddItem()
  }
  for i=0; i<NumItems; i++ {
    if ProductDB[i] = ProductID {
      break
    }
  }
  if(i == NumItems){
    // User does not exist. Add to UserDB.
    ProductDB = append(ProductDB, ProductID)
    NumItems = NumItems+1
    ALS.AddItem()
  }
  productNumber := i

  // Now add this to the Ratings matrix.
  ALS.Ratings_Matrix.Set(userNumber, productNumber, weightToAdd)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOk)
}

func GetNumUsers() (NumUsers int){
  return NumUsers
}

func GetNumItems() (NumItems int){
  return NumItems
}
