package main

import "strconv"

func CountZeroPointings(inputs []string) int {
	maxPosition := 99
	minPosition := 0
	positionCounts := minPosition + maxPosition + 1
	position := 50
	zeroCount := 0
	for _, input := range inputs {
		moveDirection := rune(input[0])
		move, _ := strconv.Atoi(input[1:])
		move = move % positionCounts
		switch moveDirection {
		case 'L':
			position -= move
			if position < minPosition {
				position = positionCounts + position
			}
			break
		case 'R':
			position += move
			if position > maxPosition {
				position = position - positionCounts
			}
			break
		}
		if position == 0 {
			zeroCount++
		}
	}
	return zeroCount
}
