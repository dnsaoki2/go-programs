package main

import (
	"fmt"
	//"math"
)

func VPytha(a, b, c int) bool {
	if  a < b && b < c && a < c && a*a+b*b == c*c {
		return true
	}
	return false
}

func find() (int, int, int) {
	b := 4
	for {
		a := ((1000*1000 - 2000*b) / (2000 - 2*b)) 
		c := 1000 - a - b
		if VPytha(a, b, c) {
			return a,b,c
		}
		b++
	}
}

func main(){
	a, b, c := find()
	fmt.Println(a*b*c)
}