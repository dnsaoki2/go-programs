package main

import "testing"

func TestDifferenceSquares(t *testing.T) {
	result1 := sumSquares(10)
	result2 := squaresSum(10)
	difference := result2-result1 
	if difference != 2640 {
		t.Error("Expected 2640, got ", difference)
	}
}
