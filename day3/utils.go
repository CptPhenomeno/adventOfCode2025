package main

import (
	"math"
	"sync"
)

func convertToInts(input string, wg *sync.WaitGroup, ch chan []int) {
	defer wg.Done()
	ints := make([]int, 0)
	for _, char := range input {
		ints = append(ints, int(char-'0'))
	}
	ch <- ints
}

func rangeOf(value, n int) []int {
	values := make([]int, 0)
	for i := 0; i < n; i++ {
		values = append(values, value)
	}
	return values
}

func replaceInSlice(slice []int, start int, newItem int) {
	for i := start; i < len(slice); i++ {
		slice[i] = newItem
	}
}

func fromIntArrayToNumber(ints []int) int {
	mul := int(math.Pow10(len(ints) - 1))
	value := 0
	for _, v := range ints {
		value += v * mul
		mul /= 10
	}
	return value
}
