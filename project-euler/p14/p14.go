package main

import (
	"fmt"
)

func collatzSequence(start int) int {
	count := 0
	for start > 1 {
		count++
		if start % 2 == 0 {
			start = start / 2
		} else {
			start = start * 3 + 1
		}
	}
	count++
	return count
}

func main() {
	fmt.Println(collatzSequence(999999))
	larger := 0
	index := 0
	for i := 2; i <= 1000000; i++{
		value := collatzSequence(i) 
		if  value > larger {
			larger = value
			index = i
		}
	}
	fmt.Println(larger, index)
}