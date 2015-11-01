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
  factors:=10
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
	UpdatePreferenceMatrix(Ratings_Matrix, Preference_Matrix, numUsers, numItems)
	UpdateConfidenceMatrix(Preference_Matrix, Confidence_Matrix, numUsers, numItems)
}

// a function to set the values for a given row
func setRow(mat *DenseMatrix, which int, row []float64) *DenseMatrix {
	if mat.Cols() != len(row) {
		fmt.Println("The row to set needs to be the same dimension as the matrix")
	}
	// iterate over columns to set the values for a selected row
	for i := 0; i < mat.Cols(); i++ {
		mat.Set(which, i, row[i])
	}
	return mat
}

// a function to set the values for a given column
func setCol(mat *DenseMatrix, which int, col []float64) *DenseMatrix {
	if mat.Rows() != len(col) {
		fmt.Println("The column to set needs to be the same dimension as the matrix")
	}
	// iterate over rows to set the values for a selected columns
	for i := 0; i < mat.Rows(); i++ {
		mat.Set(i, which, col[i])
	}
	return mat
}

func ALS(){
	lambda := 0.1
	yTy := Product(Y.Transpose(), Y)

	for i := 0; i < numUsers; i++ {
		// Identity matrix of order n X n.
		CU := Eye(numItems)
		I := Eye(numItems)

		for ii :=0; ii < numItems; ii++ {
			CU.Set(ii, ii, CM.Get(i, ii))
		}

		Middle := Difference(CU, I)
		Middle = Product(Y.Transpose(), Middle)
		Middle = Product(Middle, Y)
		ToBeInversed := yTy.Add(Middle)
		ToBeInversed := ToBeInversed.Add(Scaled(I, lambda))
		Inversed, err := ToBeInversed.Inverse()
		if(err != nil){
			log.Println("Matrix inversion failed.")
		}
		Inversed_yT := Product(Inversed, Y.Transpose())
		Inversed_yT_Cu := Product(Inversed_yT, CU)

		// Multiply this by the p(u) vector.
}
