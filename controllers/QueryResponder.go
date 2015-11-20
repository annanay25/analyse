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

  // First identify user.
  userID := req.UserID
  num := req.Num

  i:=0
  for i=0; i<NumUsers; i++ {
    if UserDB[i] = UserID {
      break
    }
  }
  userNo := i
}

func TopNRec(num int, userNo int)(recommendations []int){

  usersNumber := controllers.GetNumUsers()
  productsNumber := controllers.GetNumItems()
  // Get "num" number of top recommendations
  if(num > productsNumber){
    num = productsNumber
  }
  TopProducts := make([]int, productsNumber)
  for i:=0; i<productsNumber; i++ {
    TopProducts[i]=-1
  }

  // Calculate the vector product of x[num] and yj. And get top "num" results.
  for i in range productsNumber {
    
  }
}
