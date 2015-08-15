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

func main() {
	cont := 0
	number := 2
	for cont < 10001 {
		if isPrime(number) {
			cont++
		}
		number++
	}
	fmt.Println(number-1)
}