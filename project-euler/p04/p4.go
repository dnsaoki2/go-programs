package main

import (
	"fmt"
)

func VerifyPalin(num int) bool {
	number := num
	inver := 0
	for number != 0 {
		dig := number % 10
		inver = (inver * 10) + dig
		number = number / 10
	}
	if num == inver {
		return true
	}
	return false
}

func palin() (int, int,int) {
	var num1, num2, numRet1, numRet2 int
	var larger int = 0
	for num1 = 999; num1 >= 100; num1-- {
		for num2 = 999; num2 >= 100; num2-- {
			mul := num1 * num2
			if VerifyPalin(mul) && mul > larger {
				larger = mul
				numRet1 = num1
				numRet2 = num2
			}
		}
	}
	return numRet1, numRet2, larger
}

func main() {
	num1, num2, larger := palin()
	fmt.Println(num1,num2,larger)
}
