package main

import "testing"

func TestPalindromic(t *testing.T) {
	result := VerifyPalin(906609)
	if result != true {
		t.Error("Expected true, got ", result)
	}
}
