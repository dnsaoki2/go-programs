package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(4)
	if result != 5832 {
		t.Error("Expected 5832, got ", result)
	}
}
