package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := readInputFile()
	sum := CalculateTotalJoltage(inputs)
	genericSum := GenericCalculateTotalJoltage(inputs, 12)
	fmt.Printf("sum: %v - genericSum: %v\n", sum, genericSum)
}

func readInputFile() []string {
	file, err := os.OpenFile("./day3/code_input.txt", os.O_RDONLY, 0444)
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
