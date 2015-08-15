//NAO CORRETO
package main

import "fmt"

func fib(value int) int {
	prev, cur := 0, 1
	sum := 0

	for cur < value {
		prev, cur = cur, prev+cur
		if cur % 2 == 0 {
			sum += cur
		}
	}
	return sum
}

func main() {
	fmt.Println(fib(4000000))
}

