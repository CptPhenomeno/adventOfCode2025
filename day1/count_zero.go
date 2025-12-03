package main

import (
	"fmt"
	"strconv"
)

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

func SumZeroPositions(inputs []string) int {
	maxPosition := 99
	minPosition := 0
	positionCounts := minPosition + maxPosition + 1
	position := 50
	zeroCount := 0
	for _, input := range inputs {
		moveDirection := rune(input[0])
		move, _ := strconv.Atoi(input[1:])
		if move > positionCounts {
			zeroCount += move / positionCounts
		}
		move = move % positionCounts
		startPosition := position
		switch moveDirection {
		case 'L':
			position -= move
			if position < minPosition {
				if startPosition != 0 {
					zeroCount++
				}
				position = positionCounts + position
			} else if position == 0 {
				zeroCount++
			}
			break
		case 'R':
			position += move
			if position > maxPosition {
				if startPosition != 0 {
					zeroCount++
				}
				position = position - positionCounts
			} else if position == 0 {
				zeroCount++
			}
			break
		}
		fmt.Println("")
	}
	return zeroCount
}
