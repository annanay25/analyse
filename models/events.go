package models

type struct Event {
  userID string
  productID string
  weightToAdd int   // The idea is that the weight will get summed if the user sees the same product repeatedly.
}
