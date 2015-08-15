package main

import (
	"fmt"
)

func primeFactor(num int) int {
	factor := 2
	for num > 1 {
		if num % factor == 0 {
			num = num / factor
		} else {
			factor++
		}
	}
	return factor
}

func main() {
	fmt.Println(primeFactor(600851475143))
}
