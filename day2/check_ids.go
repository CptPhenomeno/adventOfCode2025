package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func IsInvalidId(id int) bool {
	stringId := strconv.Itoa(id)
	idLength := len(stringId)
	if idLength%2 == 1 {
		return false
	}
	firstHalf := stringId[:idLength/2]
	secondHalf := stringId[idLength/2:]
	return firstHalf == secondHalf
}

func checkIds(min int, max int, wg *sync.WaitGroup, ch chan int) int {
	defer wg.Done()
	sum := 0
	for i := min; i <= max; i++ {
		if IsInvalidId(i) {
			fmt.Printf("Invalid ID: %d\n", i)
			sum += i
		}
	}
	ch <- sum
	return sum
}

func GetInvalidIdsSum(input []string) int {
	sumValues := make(chan int, 100)
	var wg sync.WaitGroup
	for _, line := range input {
		idsRange := strings.Split(line, "-")
		minValue, _ := strconv.Atoi(idsRange[0])
		maxValue, _ := strconv.Atoi(idsRange[1])
		wg.Add(1)
		go checkIds(minValue, maxValue, &wg, sumValues)
	}
	wg.Wait()
	close(sumValues)

	totalSum := 0
	for sum := range sumValues {
		totalSum += sum
	}

	return totalSum
}
