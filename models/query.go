package models

// The userID and the number of items to be predicted. Also, list of black listed items. (An example : the seen ones)
type struct query {
  userID string
  num int
  blackListed []string
}

type struct itemScores {
  userID string
  productID string
  score int
}

type struct prediction {
  predictedResult []itemScores
}
