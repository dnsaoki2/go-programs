package main

import "testing"

func TestIsPrime(t *testing.T) {
	result := isPrime(113)
	result2 := isPrime(4)
	if result != true {
		t.Error("Expected true, got ", result)
	}
	if result2 != false {
		t.Error("Expected false, got ", result2)
	}
}