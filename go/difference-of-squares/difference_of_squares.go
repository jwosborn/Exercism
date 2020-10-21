package diffsquares

import "math"

// SquareOfSum adds every number from 0 to the integer provided and squares the sum
func SquareOfSum(a int) int {
	sum := 0
	for i := 0; i <= a; i++ {
		sum += i
	}
	return sum * sum
}

//SumOfSquares adds the square of every number from 0 to the integer provided and adds the square
func SumOfSquares(a int) int {
	sum := 0
	for i := 0; i <= a; i++ {
		sum += (i * i)
	}
	return sum
}

//Difference calculates the difference between the sum of the squares and the square of the sum of an integer.
func Difference(a int) int {
	return int(math.Abs(float64(SumOfSquares(a) - SquareOfSum(a))))
}
