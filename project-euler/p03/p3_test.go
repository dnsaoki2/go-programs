package main 

import "testing"

func TestPrimeFactor(t *testing.T) {
  result := primeFactor(13195)
  if result != 29 {
    t.Error("Expected 29, got ", result)
  }
}