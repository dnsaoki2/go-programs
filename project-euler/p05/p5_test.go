package main

import "testing"

func TestSmallest(t *testing.T) {
	result := smallest(10)
	if result != 2520 {
		t.Error("Expected 2520, got ", result)
	}
}
