package models

// The userID and the number of items to be predicted. Also, list of black listed items. (An example : the seen ones)
type struct query {
  UserID string
  Num int     // Number of recommendations required.
}

type struct prediction {
  predictedResult []itemScores
}
