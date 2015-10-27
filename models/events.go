package models

type struct Event {
  userID string 'json:"userID"'
  productID string  'json:"productID"'
  weightToAdd int   'json:"weightToAdd"'
  // The idea is that the weight will get summed if the user sees the same product repeatedly.
}
