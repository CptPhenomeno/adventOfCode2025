package main

import (
	"testing"
)

func TestCountZeroPointings_Example(t *testing.T) {
	input := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}

	expected := 3
	result := CountZeroPointings(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountZeroPointings_StartAtFifty(t *testing.T) {
	input := []string{"R10", "L10"}

	expected := 0 // Starts at 50, moves to 60, then back to 50
	result := CountZeroPointings(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountZeroPointings_SingleRotation(t *testing.T) {
	input := []string{"R50"}

	expected := 1 // Starts at 50, moves to 0
	result := CountZeroPointings(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountZeroPointings_NoZeroCrossing(t *testing.T) {
	input := []string{"R10", "R10", "R10"}

	expected := 0 // Starts at 50, moves to 60, 70, 80
	result := CountZeroPointings(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountZeroPointings_MultipleZeroCrossings(t *testing.T) {
	input := []string{"R50", "L100", "R100"}

	expected := 3 // 50 -> 0 -> 0 -> 0
	result := CountZeroPointings(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountZeroPointings_EmptyInput(t *testing.T) {
	input := []string{}

	expected := 0 // Dial starts at 50, no moves to reach 0
	result := CountZeroPointings(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCountZeroPointings_LeftToZero(t *testing.T) {
	input := []string{"L50"}

	expected := 1 // Starts at 50, moves left to 0
	result := CountZeroPointings(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Part 2 Tests

// The function signature now only requires the slice of moves.
// It assumes the dial always initializes at position 50.
func TestSumZeroPositions(t *testing.T) {
	tests := []struct {
		name     string
		moves    []string
		expected int
	}{
		{
			name: "AoC Example: Full Sequence",
			moves: []string{
				"L68", "L30", "R48", "L5", "R60",
				"L55", "L1", "L99", "R14", "L82",
			},
			expected: 6,
		},
		{
			name:     "AoC Warning: Massive Rotation",
			moves:    []string{"R1000"},
			expected: 10,
			// Logic: Starts 50. 50 + 1000 = 1050.
			// 1000 / 100 = 10 full loops.
			// Passes 0 ten times and lands back on 50.
		},
		{
			name:     "Directly to 0 from Right",
			moves:    []string{"R50"},
			expected: 1, // 50 -> 100 (which is 0). Counts as 1.
		},
		{
			name:     "Directly to 0 from Left",
			moves:    []string{"L50"},
			expected: 1, // 50 -> 0. Counts as 1.
		},
		{
			name:     "Pass 0 from Right (Single Move)",
			moves:    []string{"R60"},
			expected: 1, // 50 -> 110 (Pos 10). Passed 0.
		},
		{
			name:     "Pass 0 from Left (Single Move)",
			moves:    []string{"L60"},
			expected: 1, // 50 -> -10 (Pos 90). Passed 0.
		},
		{
			name:     "Move Away from 0 (Left)",
			moves:    []string{"L50", "L5"},
			expected: 1,
			// 1. Start 50 -> L50 lands on 0 (Hit #1).
			// 2. Start 0  -> L5 lands on 95 (Hit #0).
			// Total: 1.
		},
		{
			name:     "Move Away from 0 (Right)",
			moves:    []string{"L50", "R5"},
			expected: 1,
			// 1. Start 50 -> L50 lands on 0 (Hit #1).
			// 2. Start 0  -> R5 lands on 5 (Hit #0).
			// Total: 1.
		},
		{
			name:     "Full Circle from 0",
			moves:    []string{"L50", "R100"},
			expected: 2,
			// 1. Start 50 -> L50 lands on 0 (Hit #1).
			// 2. Start 0  -> R100 lands on 0 (Hit #2).
			// Total: 2.
		},
		{
			name:     "Multiple Rotations Landing on 0",
			moves:    []string{"R250"},
			expected: 3,
			// Start 50.
			// 50 + 250 = 300.
			// 300 is exactly 0 mod 100.
			// It passes 100 (1), 200 (2), lands on 300 (3).
		},
		{
			name:     "Zig Zag Crossing 0",
			moves:    []string{"R60", "L20"},
			expected: 2,
			// 1. Start 50 -> R60 -> Pos 10 (Total 110). passed 0. (Hit #1)
			// 2. Start 10 -> L20 -> Pos 90 (Total -10). passed 0. (Hit #2)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: We no longer pass startPos
			got := SumZeroPositions(tt.moves)
			if got != tt.expected {
				t.Errorf("CalculateZeroHits() = %v, want %v", got, tt.expected)
			}
		})
	}
}
