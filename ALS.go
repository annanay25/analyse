package ALS

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/skelterjohn/go.matrix"
)

// Creating the matrices.
func MakeXY(numUsers int, numItems int, max_rating float64) (X,Y *DenseMatrix) {
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

func MakeRatingsMatrix(numUsers int, numItems int)(Ratings_Matrix *DenseMatrix){
  Ratings_data := make([]float64, numUsers*numItems)
  for i := 0; i < len(Ratings_data); i++ {
			Ratings_Matrix[i]=0
	}
	return Ratings_Matrix := MakeDenseMatrix(X_data, numUsers, numItems)
}

func MakePreferenceMatrix(numUsers int, numItems int)(Preference_Matrix *DenseMatrix){
	Preference_data := make([]float64, numUsers*numItems)
	for i := 0; i < len(Preference_data); i++ {
			Preference_data[i]=0
	}
	Preference_Matrix := MakeDenseMatrix(Preference_data, numUsers, numItems)
}

func MakeConfidenceMatrix(numUsers int, numItems int)(Confidence_Matrix *DenseMatrix){
	Confidence_data := make([]float64, numUsers*numItems)
	for i := 0; i < len(Confidence_data); i++ {
			Confidence_data[i]=0
	}
	Confidence_Matrix := MakeDenseMatrix(Confidence_data, numUsers, numItems)
}

// Updates the preference matrix. Function call is from QueryController.
func UpdatePreferenceMatrix(RM *DenseMatrix,PM *DenseMatrix,numUsers int, numItems int)(Preference_Matrix *DenseMatrix){
	// ROWS: numUsers, COLS: numItems.
	for i := 0; i < numUsers; i++ {
		for j := 0; j < numItems ; j++{
			if RM.Get(i, j) > 0 {
				PM.Set(i, j, 1)
			} else {
				PM.Set(i, j, 0)
			}
		}
	}
	return PM
}

// Creates the confidence matrix.
func UpdateConfidenceMatrix(PM *DenseMatrix,CM *DenseMatrix,numUsers int, numItems int)(Confidence_Matrix *DenseMatrix){
	alpha := 40

	for i := 0; i < numUsers; i++ {
		for j := 0; j < numItems ; j++{
			value := 1 + alpha * PM.Get(i, j)
			CM.Set(i, j, value)
		}
	}
	return CM
}

func Prepare(){
	// Make the User and Item matrices.
  numUsers:= controllers.GetUserInfo()
  numItems:= controllers.GetProductInfo()
  Ratings_Matrix:=MakeRatingsMatrix(numUsers, numItems)
	Preference_Matrix := MakePreferenceMatrix(Ratings_Matrix, numUsers, numItems)
	Confidence_Matrix := MakeConfidenceMatrix(Preference_Matrix, numUsers, numItems)
  // We need max_rating from this but we cannot get that until training is done.
  X, Y:= MakeXY(numUsers, numItems)
}

func Train(){
	

}
