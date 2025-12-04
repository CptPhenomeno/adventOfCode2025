package main

import (
	"testing"
)

func TestGenericFindMaxJoltageForBank(t *testing.T) {
	tests := []struct {
		name      string
		bank      []int // Input is a slice of joltage ratings (1-9)
		batteries int
		expected  int
	}{
		// --- Examples from the Problem Description ---
		{
			name:      "Example 1: Largest first two digits (98)",
			bank:      []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			batteries: 12,
			expected:  987654321111,
		},
		{
			name:      "Example 2: Max at ends (89)",
			bank:      []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			batteries: 12,
			expected:  811111111119,
		},
		{
			name:      "Example 3: Max at last two (78)",
			bank:      []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			batteries: 12,
			expected:  434234234278,
		},
		{
			name:      "Example 4: Max non-sequential (92)",
			bank:      []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			batteries: 12,
			expected:  888911112111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenericFindMaxJoltageForBank(tt.bank, tt.batteries); got != tt.expected {
				t.Errorf("FindMaxJoltageForBank(%v) = %d, want %d", tt.bank, got, tt.expected)
			}
		})
	}
}

func TestGenericCalculateTotalJoltage(t *testing.T) {
	// Input is a slice of strings, where each string is a battery bank.
	banks := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}

	expectedTotal := 3121910778619

	t.Run("AoC Example: Total Output Joltage", func(t *testing.T) {
		if got := GenericCalculateTotalJoltage(banks, 12); got != expectedTotal {
			t.Errorf("CalculateTotalJoltage(%v) = %d, want %d", banks, got, expectedTotal)
		}
	})
}
