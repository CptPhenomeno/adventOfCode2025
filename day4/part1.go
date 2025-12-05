package main

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

func CollectRollPaper(matrix Matrix) int {
	result := IdentifyValidRollPaper(matrix)
	count := 0
	for r, row := range result {
		for c, cell := range row {
			if cell == 2 {
				count++
				matrix[r][c] = 0
			}
		}
	}
	return count
}

func CountValidRollPaper(matrix Matrix) int {
	return CollectRollPaper(matrix)
}
