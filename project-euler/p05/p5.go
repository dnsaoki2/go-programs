package main

import (
	"fmt"
)

func smallest(lim int) int {
	var num int = 1
	var cont int = 0
	var i int
	for {
		for i = 1; i <= lim; i++ {
			if num % i != 0 {
				num++
			} else {
				cont++
			} 
		}
		if cont >= lim {
			return num
		} else{
			cont = 0
		}
	}
	return 0
}

func main() {
	fmt.Println(smallest(20))
}
