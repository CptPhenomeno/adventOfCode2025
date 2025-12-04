package main

import (
	"sync"
)

func FindMaxJoltageForBank(bank []int) int {
	tens := -1
	units := -1
	size := len(bank)
	for _, joltage := range bank[:size-1] {
		if joltage > tens {
			tens = joltage
			if units != -1 {
				units = -1
			}
		} else if joltage > units {
			units = joltage
		}
	}
	if bank[size-1] > units {
		units = bank[size-1]
	}
	return (tens * 10) + units
}

func CalculateTotalJoltage(banks []string) int {
	wg := sync.WaitGroup{}
	ch := make(chan []int, len(banks))
	for _, bank := range banks {
		wg.Add(1)
		go convertToInts(bank, &wg, ch)
	}
	wg.Wait()
	close(ch)

	intBanks := make([][]int, 0)
	for ints := range ch {
		intBanks = append(intBanks, ints)
	}

	sum := 0
	for _, bank := range intBanks {
		sum += FindMaxJoltageForBank(bank)
	}
	return sum
}
