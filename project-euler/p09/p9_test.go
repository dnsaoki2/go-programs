package main

import "testing"

func TestPalindromic(t *testing.T) {
	result := VPytha(3, 4, 5)
	result2 := VPytha(9, 5, 5)
	if result != true {
		t.Error("Expected true, got ", result)
	}
	if result2 != false {
		t.Error("Expected false, got ", result)
	}
}
