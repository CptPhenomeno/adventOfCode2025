package main

import (
	"testing"
)

// The main function signature being tested is assumed to be:
// func GetInvalidIdsSum(input string) int
// Which calculates the sum of all invalid IDs across all ranges.

// --- Core Logic Test ---
// TestIsInvalidID checks the fundamental rule: an ID is invalid if it is made
// only of some sequence of digits repeated twice (e.g., 55, 6464).
func TestIsInvalidID(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want bool
	}{
		// --- True Cases (Valid Invalid IDs) ---
		{"2-digit repeat: 55", 55, true},
		{"4-digit repeat: 6464", 6464, true},
		{"6-digit repeat: 123123", 123123, true},
		{"4-digit repeat: 1010", 1010, true},
		{"6-digit repeat: 222222 (222 repeated)", 222222, true},
		{"Large 10-digit repeat: 1188511885", 1188511885, true},
		{"Repeated zero sequence: 00", 0, false},     // IDs have no leading zeroes; 0 isn't '0' repeated twice in a valid way (assuming min 1-digit repeat)
		{"Repeated zero sequence: 0", 0, false},      // Single digit is never a repeat
		{"Repeated zero sequence: 0101", 101, false}, // Per rule: 0101 isn't an ID at all (but 101 is valid and ignored)

		// --- False Cases (Valid IDs to be ignored) ---
		{"Odd length: 123", 123, false},
		{"Odd length symmetric: 101", 101, false},
		{"Even length but no repeat: 1234", 1234, false},
		{"Even length with different halves: 1241", 1241, false},
		{"Triple repeat: 565656 (should be false)", 565656, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInvalidId(tt.id); got != tt.want {
				t.Errorf("IsInvalidId(%d) = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}

// -----------------------------------------------------------------------

// --- Integration Tests (GetInvalidIdsSum) ---
// These tests check the full process: parsing the ranges, iterating through the numbers,
// checking for invalid IDs, and summing the result.

func TestGetInvalidIdsSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []string // Changed to string array
		expected int
	}{
		{
			name: "AoC Example: Full Input",
			// The original long comma-separated string is now split into an array of ranges.
			input: []string{
				"11-22", "95-115", "998-1012", "1188511880-1188511890",
				"222220-222224", "1698522-1698528", "446443-446449",
				"38593856-38593862", "565653-565659", "824824821-824824827",
				"2121212118-2121212124",
			},
			expected: 1227775554,
		},
		{
			name: "Range 1: Simple 2-Digit Repeats",
			// Invalid IDs are 11 and 22. Sum = 33.
			input:    []string{"11-22"},
			expected: 33,
		},
		{
			name: "Range 2: Single 4-Digit Repeat",
			// Range 998-1012 contains 1010. Sum = 1010.
			input:    []string{"998-1012"},
			expected: 1010,
		},
		{
			name: "Range 3: Near Boundary Check",
			// Range 222220-222224 contains 222222 (222 repeated). Sum = 222222.
			input:    []string{"222220-222224"},
			expected: 222222,
		},
		{
			name: "Range 4: Multiple, Non-Consecutive Invalid IDs",
			// Range 60-120
			// Invalid IDs: 66, 77, 88, 99, 1010, 1111 (11 repeated).
			// Sum: 66 + 77 + 88 + 99 + 1010 + 1111 = 2551
			input:    []string{"60-1200"},
			expected: 2451,
		},
		{
			name: "Range 5: Zero Invalid IDs",
			// Range 1698522-1698528 contains no invalid IDs per example.
			input:    []string{"1698522-1698528"},
			expected: 0,
		},
		{
			name: "Range 6: Multiple Ranges, Simple Sum",
			// 10-15: Invalid: 11.
			// 30-45: Invalid: 33, 44.
			// 990-1015: Invalid: 1010.
			// Sum: 11 + 33 + 44 + 1010 = 1098
			input:    []string{"10-15", "30-45", "990-1015"},
			expected: 1098,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// The call to SolvePart1 now passes a slice of strings
			if got := GetInvalidIdsSum(tt.input); got != tt.expected {
				t.Errorf("SolvePart1(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
