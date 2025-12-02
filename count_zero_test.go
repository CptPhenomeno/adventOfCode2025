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
