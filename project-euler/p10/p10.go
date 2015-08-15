package main 

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 { 
		return false
	}
	if n == 2{
		return true
	}
	if n % 2 == 0{
		return false
	}
	if n < 9 {
		return true
	} 
	if n % 3 == 0 {
		return false
	}
	max := math.Sqrt(float64(n))+1
	for i := 5; i <= int(max); i += 6 {
		if n % i == 0 {
			return false
		}
		if n % (i + 2) == 0 {
			return false
		}
	}
	return true
}

func sumPrime(lim int) int {
	value := 2
	sum := 0

	for value < lim {
		if isPrime(value) {
			sum += value
		}
		value++
	}
	return sum
}

func main() {
	fmt.Println(sumPrime(2000000))
}