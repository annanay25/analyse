package ALS

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/skelterjohn/go.matrix"
)


func makeXY(numUsers int, numItems int, max_rating float64) (X,Y *DenseMatrix) {
  rand.Seed(int64(seed))
  factors:=5
  X_data := make([]float64, numUsers*factors)
  Y_data := make([]float64, numItems*factors)

  for i := 0; i < len(X_data); i++ {
		X_data[i] = max_rating * rand.Float64()
	}
	for j := 0; j < len(Y_data); j++ {
		Y_data[j] = max_rating * rand.Float64()
	}
	X = MakeDenseMatrix(X_data, factors, numUsers)
	Y = MakeDenseMatrix(Y_data, factors, numItems)

  return
}

// Creates the main ratings matrix based on which ALS will operate.
func MakeRatingsMatrix(numUsers int, numItems int)(Ratings_matrix *DenseMatrix){
  Ratings_data := make([]float64, numUsers*numItems)
  for i := 0; i < len(Ratings_data); i++ {
			Ratings_matrix[i]=0
	}
	return Ratings_matrix := MakeDenseMatrix(X_data, numUsers, numItems)
}

// Creates the preference matrix.
func MakePreferenceMatrix(RM *DenseMatrix,numUsers int, numItems int)(Preference_matrix *DenseMatrix){
	Preference_data := make([]float64, numUsers*numItems)
	for i := 0; i < len(Ratings_data); i++ {
			if RM[i] > 0 {
				Preference_data[i]=1
			} else {
				Preference_data[i]=0
			}
	}
	return Preference_Matrix := MakeDenseMatrix(Preference_data, numUsers, numItems)
}

// Creates the confidence matrix.
func MakeConfidenceMatrix(RM *DenseMatrix,numUsers int, numItems int)(Confidence_Matrix *DenseMatrix){
	alpha := 40
	Confidence_data := make([]float64, numUsers*numItems)
	for i := 0; i < len(Ratings_data); i++ {
			row := i/numItems
			col := i%numItems

			Confidence_data[i] = 1 + alpha * RM[row][col]
	}
	return Confidence_Matrix := MakeDenseMatrix(Confidence_data, numUsers, numItems)
}


func Train(){

  // Make the User and Item matrices.
  numUsers:= controllers.GetUserInfo()
  numItems:= controllers.GetProductInfo()
  ratings:=MakeRatingsMatrix(numUsers, numItems)
  // We need max_rating from this but we cannot get that until training is done.
  X, Y:= makeXY(numUsers, numItems)
}
