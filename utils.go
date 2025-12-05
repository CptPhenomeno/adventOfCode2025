package adventOfCode2025

import (
	"math"
	"sync"
)

func ConvertToInts(input string, wg *sync.WaitGroup, ch chan []int) {
	defer wg.Done()
	ints := make([]int, 0)
	for _, char := range input {
		ints = append(ints, int(char-'0'))
	}
	ch <- ints
}

func ConvertToIntArray[T int | uint8](input string) []T {
	ints := make([]T, 0)
	for _, char := range input {
		ints = append(ints, T(char-'0'))
	}
	return ints
}

func RangeOf(value, n int) []int {
	values := make([]int, 0)
	for i := 0; i < n; i++ {
		values = append(values, value)
	}
	return values
}

func ReplaceInSlice(slice []int, start int, newItem int) {
	for i := start; i < len(slice); i++ {
		slice[i] = newItem
	}
}

func FromIntArrayToNumber(ints []int) int {
	mul := int(math.Pow10(len(ints) - 1))
	value := 0
	for _, v := range ints {
		value += v * mul
		mul /= 10
	}
	return value
}
