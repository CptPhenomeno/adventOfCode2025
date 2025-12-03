package main

import (
	"reflect"
	"testing"
)

func TestGetDivisors(t *testing.T) {
	tests := []struct {
		name  string
		input int
		// Divisors should be returned in sorted order (ascending)
		expected []int
	}{
		// Edge Case
		{"Input 1", 1, []int{1}},

		// Prime Numbers
		{"Prime 7", 7, []int{1, 7}},
		{"Prime 13", 13, []int{1, 13}},

		// Composite Numbers
		{"Composite 10", 10, []int{1, 2, 5, 10}},
		{"Composite 12", 12, []int{1, 2, 3, 4, 6, 12}},
		{"Composite 30", 30, []int{1, 2, 3, 5, 6, 10, 15, 30}},

		// Perfect Square
		{"Perfect Square 36", 36, []int{1, 2, 3, 4, 6, 9, 12, 18, 36}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Assumes your function is named FindDivisors
			got := GetDivisors(tt.input)

			// Use reflect.DeepEqual to compare the slices
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindDivisors(%divisorCache) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGetInvalidMultiIdsSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []string // []string{"min-max", ...}
		expected int
	}{
		{
			name: "AoC Part 2 Example: Full Input",
			// Input ranges remain the same as Part 1
			input: []string{
				"11-22", "95-115", "998-1012", "1188511880-1188511890",
				"222220-222224", "1698522-1698528", "446443-446449",
				"38593856-38593862", "565653-565659", "824824821-824824827",
				"2121212118-2121212124",
			},
			// The new expected sum for Part 2
			expected: 4174379265,
		},
		{
			name: "Range 1: Simple 2-Digit Repeats (Unchanged)",
			// Invalid IDs: 11, 22. Sum = 33.
			input:    []string{"11-22"},
			expected: 33,
		},
		{
			name: "Range 2: Odd Length Repeats Introduced",
			// 95-115 now has 99 (2x) and 111 (3x). Sum = 99 + 111 = 210.
			input:    []string{"95-115"},
			expected: 210,
		},
		{
			name: "Range 3: 999 Introduced",
			// 998-1012 now has 999 (3x) and 1010 (2x). Sum = 999 + 1010 = 2009.
			input:    []string{"998-1012"},
			expected: 2009,
		},
		{
			name: "Range 4: 565656 Introduced",
			// 565653-565659 now has 565656 (56 repeated three times). Sum = 565656.
			input:    []string{"565653-565659"},
			expected: 565656,
		},
		{
			name: "Range 5: Very long repeat introduced",
			// 2121212118-2121212124 now has 2121212121 (21 repeated five times). Sum = 2121212121.
			input:    []string{"2121212118-2121212124"},
			expected: 2121212121,
		},
		{
			name: "Range 6: Combined New Logic",
			// 1-100: Invalid IDs: 11, 22, ..., 99. Sum = 495.
			// 100-200: Invalid IDs: 111. Sum = 111.
			// Total Sum: 495 + 111 = 606.
			input:    []string{"1-100", "100-200"},
			expected: 606,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInvalidMultiIdsSum(tt.input); got != tt.expected {
				t.Errorf("GetInvalidMultiIdsSum(%v) = %divisorCache, want %divisorCache", tt.input, got, tt.expected)
			}
		})
	}
}
