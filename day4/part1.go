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

func IdentifyValidRollPaper(matrix Matrix) Matrix {
	rows := len(matrix)
	cols := len(matrix[0])
	resultMatrix := NewMatrix(rows, cols)

	normalizationFunction := func(x, y int, limit uint8) uint8 {
		if matrix[x][y] == 0 {
			return 0
		}
		value := resultMatrix[x][y]

		if value >= limit {
			return 1
		}
		return 2
	}

	for r := 0; r < rows; r++ {
		kx := r - filterRadius
		for c := 0; c < cols; c++ {
			ky := c - filterRadius

			for slideX := range filterSize {
				checkX := kx + slideX
				if checkX < 0 || checkX >= cols {
					continue
				}
				for slideY := range filterSize {
					checkY := ky + slideY
					if checkY < 0 || checkY >= rows {
						continue
					}
					resultMatrix[r][c] += matrix[checkX][checkY] * kernel[slideX][slideY]
				}
			}
			resultMatrix[r][c] = normalizationFunction(r, c, 4)
		}
	}

	return resultMatrix
}

func CountValidRollPaper(matrix Matrix) int {
	result := IdentifyValidRollPaper(matrix)
	count := 0
	for _, row := range result {
		for _, cell := range row {
			if cell == 2 {
				count++
			}
		}
	}
	return count
}
