package main

import "testing"

func TestSumPrime(t *testing.T) {
	result := isPrime(13)
	result2 := isPrime(14)
	result3 :=  sumPrime(10)
	if result != true {
		t.Error("Expected true, got ", result)
	}
	if result2 != false {
		t.Error("Expected false, got ", result2)
	}
	if result3 != 17 {
		t.Error("Expected 17, got ", result3)
	}
}
