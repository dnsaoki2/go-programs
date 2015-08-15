package main

import "testing"

func TestEvenFib(t *testing.T) {
	result := fib(89)
	if (result != 44) {
		t.Error("Expected 995 and 583, got ", result)
	}
}
