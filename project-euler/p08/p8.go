package main 

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func sum(n int) float64 {
	larger := 0.0
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i <= len(data) - n; i++ {
		mul := 1.0
		for j := 0; j < n; j++ {
			value, _ := strconv.Atoi(string(data[i+j])) 
			mul *= float64(value)
		}
		if mul > larger {
			larger = mul
		}
	} 
	return larger
}

func main() {
	fmt.Println(sum(13))
}