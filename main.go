package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := readInputFile()
	count := CountZeroPointings(inputs)
	sum := SumZeroPositions(inputs)
	fmt.Printf("count: %v sum: %v\n", count, sum)
}

func readInputFile() []string {
	file, err := os.OpenFile("code_input.txt", os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	inputs := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	return inputs
}
