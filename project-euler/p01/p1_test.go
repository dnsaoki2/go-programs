package main 

import "testing"

func TestSumMult(t *testing.T) {
  result := sum(10)
  if result != 23 {
    t.Error("Expected 23, got ", result)
  }
}