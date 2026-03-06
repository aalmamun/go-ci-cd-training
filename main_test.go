package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct{
		a int
		b int
		expected int
	}{
		{1,2,3},
		{5,5,10},
		{-1,1,0},
	}

	for _, tt := range tests {
		result := add(tt.a, tt.b)

		if result != tt.expected {
			t.Errorf("add(%d,%d) expected %d got %d",
				tt.a, tt.b, tt.expected, result)
		}
	}
}
