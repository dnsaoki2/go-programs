//Soma do multiplos de 3 e 5 abaixo de 1000
package main

import (
	"fmt"
)

func sum(lim int) int {
	sum := 0
	for value := 1; value < lim; value++ {
		if (value % 3 == 0) || (value % 5 == 0) {
			sum += value
		}
	}
	return sum

}

func main() {
	fmt.Println(sum(1000))
}
