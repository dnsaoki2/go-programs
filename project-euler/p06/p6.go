//Problema 6
package main

import "fmt"

func sumSquares(quant int) float64 {
	sum := 0.0
	for i := 0; i <= quant; i++ {
		sum += float64(i * i)
	}
	return sum
}

func squaresSum(quant int) float64 {
	sum := 0.0
	for i := 0; i <= quant; i++ {
		sum += float64(i)
	}
	return sum * sum

}

func main() {
	SumSquares := sumSquares(100)
	SquaresSum := squaresSum(100)
	fmt.Println("Sum of the Squares: ", SumSquares)
	fmt.Println("Squares of the sum: ", SquaresSum)
	fmt.Println("Difference: ", SquaresSum-SumSquares)
}
