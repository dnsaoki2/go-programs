package main

import "testing"

func TestCollatz(t *testing.T) {
	result := collatzSequence(13)
	if result != 10 {
		t.Error("Expected 10, got", result)
	}
}