package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	inputs := readInputFile()
	sum := GetInvalidIdsSum(inputs)
	sumMulti := GetInvalidMultiIdsSum(inputs)
	fmt.Printf("sum: %v - sumMulti: %v\n", sum, sumMulti)
}

func readInputFile() []string {
	file, err := os.OpenFile("./day2/code_input.txt", os.O_RDONLY, 0444)
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
	scanner.Split(scanCommas)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	return inputs
}

func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
