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

func TestTableSum(t *testing.T) {
	tests := []struct {
		input1   int
		input2   int
		expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
	}
	for _, test := range tests {
		if output := Sum(test.input1, test.input2); output != test.expected {
			t.Error("Sum was incorrect, got:", test.expected)
		}
	}
}
