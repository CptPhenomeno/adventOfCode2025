package main

import (
	"adventOfCode2025"
	"sync"
)

func checkDigits(digits []int, index int, joltage int) {
	if index == len(digits)-1 {
		if joltage > digits[index] {
			digits[index] = joltage
		}
		return
	}

	if joltage > digits[index] {
		digits[index] = joltage
		adventOfCode2025.ReplaceInSlice(digits, index+1, -1)
		return
	}
	checkDigits(digits, index+1, joltage)
}

func GenericFindMaxJoltageForBank(bank []int, batteries int) int {
	digits := adventOfCode2025.RangeOf(-1, batteries)
	limit := len(bank) - batteries
	for _, joltage := range bank[:limit] {
		checkDigits(digits, 0, joltage)
	}
	for i, joltage := range bank[limit:] {
		checkDigits(digits, i, joltage)
	}

	return adventOfCode2025.FromIntArrayToNumber(digits)
}

func GenericCalculateTotalJoltage(banks []string, batteries int) int {
	wg := sync.WaitGroup{}
	ch := make(chan []int, len(banks))
	for _, bank := range banks {
		wg.Add(1)
		go adventOfCode2025.ConvertToInts(bank, &wg, ch)
	}
	wg.Wait()
	close(ch)

	intBanks := make([][]int, 0)
	for ints := range ch {
		intBanks = append(intBanks, ints)
	}

	sum := 0
	for _, bank := range intBanks {
		sum += GenericFindMaxJoltageForBank(bank, batteries)
	}
	return sum
}
