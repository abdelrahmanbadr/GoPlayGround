package main

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	total := Sum(1, 2)
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 3)
	}
}

func Test_SumSlice(t *testing.T) {
	sum, _ := SumSlice(5)
	if sum != 15 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 3)
	}
}
