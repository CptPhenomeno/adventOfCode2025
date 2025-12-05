package main

import (
	"adventOfCode2025"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := readInputFile()
	result := CollectRollPaper(matrix)
	fmt.Printf("result: %v\n", result)

	matrix = readInputFile()
	total := CollectAllRollPaper(matrix)
	fmt.Printf("total: %v\n", total)
}

func readInputFile() Matrix {
	file, err := os.OpenFile("./day4/code_input.txt", os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	inputs := make(Matrix, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ".", "0")
		line = strings.ReplaceAll(line, "@", "1")
		inputs = append(inputs, adventOfCode2025.ConvertToIntArray[uint8](line))
	}

	return inputs
}
