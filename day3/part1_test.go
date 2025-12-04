package main

import (
	"testing"
)

func TestFindMaxJoltageForBank(t *testing.T) {
	tests := []struct {
		name     string
		bank     []int // Input is a slice of joltage ratings (1-9)
		expected int   // The maximum possible two-digit joltage
	}{
		// --- Examples from the Problem Description ---
		{
			name:     "Example 1: Largest first two digits (98)",
			bank:     []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			expected: 98, // The pair (9, 8) forms the max joltage.
		},
		{
			name:     "Example 2: Max at ends (89)",
			bank:     []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			expected: 89, // The pair (8, 9) forms the max joltage.
		},
		{
			name:     "Example 3: Max at last two (78)",
			bank:     []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			expected: 78, // The pair (7, 8) forms the max joltage.
		},
		{
			name:     "Example 4: Max non-sequential (92)",
			bank:     []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			expected: 92, // The pair (9, 2) forms the max joltage.
		},

		// --- Edge Cases and Logic Checks ---
		{
			name:     "Case 5: Max is 99 (Self-contained)",
			bank:     []int{1, 3, 5, 9, 9, 8, 7},
			expected: 99, // The pair (9, 9) forms the max joltage.
		},
		{
			name:     "Case 6: Max is 99 (Wide separation)",
			bank:     []int{9, 1, 1, 1, 1, 1, 1, 9},
			expected: 99, // The pair (9, 9) forms the max joltage.
		},
		{
			name:     "Case 7: Lowest possible (11)",
			bank:     []int{1, 1, 1, 1, 1},
			expected: 11,
		},
		{
			name:     "Case 8: Max pair is composed of non-max digits (89 vs 91)",
			bank:     []int{1, 8, 9, 1},
			expected: 91, // Pairs: (18, 19, 11, 89, 81, 91). Max is 91, but 89 must be tested. Wait: Pairs: (1,8), (1,9), (1,1), (8,9), (8,1), (9,1). Max is 91.
		},
		{
			name:     "Case 9: Correctly prioritizing first digit (91 vs 89)",
			bank:     []int{8, 9, 1},
			expected: 91, // Pairs: (89, 81, 91). Max is 91.
		},
		{
			name:     "Case 10: Shortest possible input",
			bank:     []int{4, 7},
			expected: 47,
		},
		{
			name:     "Case 11: Single highest digit followed by lowest (91)",
			bank:     []int{9, 1, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 98, // Pairs include 91, 98. Max is 98.
		},
		{
			name:     "Case 12: No 9s (Max 88)",
			bank:     []int{1, 2, 3, 8, 1, 5, 8, 1},
			expected: 88, // The pair (8, 8) forms the max joltage.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxJoltageForBank(tt.bank); got != tt.expected {
				t.Errorf("FindMaxJoltageForBank(%v) = %d, want %d", tt.bank, got, tt.expected)
			}
		})
	}
}

func TestCalculateTotalJoltage(t *testing.T) {
	// Input is a slice of strings, where each string is a battery bank.
	banks := []string{
		"987654321111111", // Max: 98
		"811111111111119", // Max: 89
		"234234234234278", // Max: 78
		"818181911112111", // Max: 92
	}

	// Sum: 98 + 89 + 78 + 92 = 357
	expectedTotal := 357

	t.Run("AoC Example: Total Output Joltage", func(t *testing.T) {
		if got := CalculateTotalJoltage(banks); got != expectedTotal {
			t.Errorf("CalculateTotalJoltage(%v) = %d, want %d", banks, got, expectedTotal)
		}
	})

	t.Run("Edge Case: Small Banks Sum", func(t *testing.T) {
		smallBanks := []string{
			"19", // Max: 19
			"91", // Max: 91
			"44", // Max: 44
		}
		// Sum: 19 + 91 + 44 = 154
		expectedSum := 154
		if got := CalculateTotalJoltage(smallBanks); got != expectedSum {
			t.Errorf("CalculateTotalJoltage(%v) = %d, want %d", smallBanks, got, expectedSum)
		}
	})

	t.Run("Edge Case: Bank with Max 99", func(t *testing.T) {
		bankWithMax99 := []string{
			"9987", // Max: 99
			"8765", // Max: 87
		}
		// Sum: 99 + 87 = 186
		expectedSum := 186
		if got := CalculateTotalJoltage(bankWithMax99); got != expectedSum {
			t.Errorf("CalculateTotalJoltage(%v) = %d, want %d", bankWithMax99, got, expectedSum)
		}
	})
}
