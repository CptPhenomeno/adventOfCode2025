package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rangeSet, ingredients := readInputFile()
	result := CountFreshIngredients(ingredients, rangeSet)
	fmt.Printf("result: %v\n", result)
}

func readInputFile() (*RangeSet, []int) {
	file, err := os.OpenFile("./day5/code_input.txt", os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	rangeSet := NewRangeSet()
	ingredients := make([]int, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		minVal, _ := strconv.Atoi(split[0])
		maxVal, _ := strconv.Atoi(split[1])
		rg := NewRange(minVal, maxVal)
		rangeSet.Add(*rg)
	}

	for scanner.Scan() {
		line := scanner.Text()
		ingredient, _ := strconv.Atoi(line)
		ingredients = append(ingredients, ingredient)
	}

	return rangeSet, ingredients
}
