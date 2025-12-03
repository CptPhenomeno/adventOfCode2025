package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type DivisorCache struct {
	Mutex sync.Mutex
	Cache map[int][]int
}

type CheckResult struct {
	IsInvalid bool
	Value     int
}

var divisorCache DivisorCache = DivisorCache{Cache: make(map[int][]int)}

func GetDivisors(n int) []int {
	defer divisorCache.Mutex.Unlock()
	divisorCache.Mutex.Lock()
	if divisors, ok := divisorCache.Cache[n]; ok {
		return divisors
	}
	divisors := findDivisors(n)
	divisorCache.Cache[n] = divisors
	return divisors
}

func findDivisors(n int) []int {
	divisors := make([]int, 0)
	stopIndex := int(math.Floor(math.Sqrt(float64(n))))
	for i := 1; i <= stopIndex; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	sort.Ints(divisors)
	return divisors
}

func checkForInvalidId(id string, windowSize int) CheckResult {
	if windowSize == len(id) {
		checkResult := CheckResult{IsInvalid: false, Value: -1}
		return checkResult
	}

	substring := id[:windowSize]
	for i := windowSize; i <= len(id)-windowSize; i = i + windowSize {
		value := id[i : i+windowSize]
		if substring != value {
			checkResult := CheckResult{IsInvalid: false, Value: -1}
			return checkResult
		}
	}
	retVal, _ := strconv.Atoi(id)
	checkResult := CheckResult{IsInvalid: true, Value: retVal}
	return checkResult
}

func checkAllPossibleInvalidIds(id int) []int {
	idString := strconv.Itoa(id)
	valueMap := make(map[int]bool, 0)
	for _, d := range GetDivisors(len(idString)) {
		result := checkForInvalidId(idString, d)
		if result.IsInvalid {
			valueMap[result.Value] = true
		}
	}
	values := make([]int, 0)
	for k, _ := range valueMap {
		values = append(values, k)
	}
	return values
}

func checkRange(minVal int, maxVal int) []int {
	values := make([]int, 0)
	for i := minVal; i <= maxVal; i++ {
		values = append(values, checkAllPossibleInvalidIds(i)...)
	}
	return values
}

func GetInvalidMultiIdsSum(ranges []string) int {
	values := make([]int, 0)
	for _, s := range ranges {
		split := strings.Split(s, "-")
		minVal, _ := strconv.Atoi(split[0])
		maxVal, _ := strconv.Atoi(split[1])
		values = append(values, checkRange(minVal, maxVal)...)
	}

	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
