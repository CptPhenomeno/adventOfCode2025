package main

import "fmt"

type Matrix [][]uint8

func NewMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]uint8, cols)
	}
	return matrix
}

func (m Matrix) String() string {
	result := ""
	for _, row := range m {
		for _, cell := range row {
			result += fmt.Sprintf("%d ", cell)
		}
		result += "\n"
	}
	return result
}

type Filter [][]uint8

const filterSize = 3
const filterRadius = filterSize / 2

var kernel = Filter{
	{1, 1, 1},
	{1, 0, 1},
	{1, 1, 1},
}
