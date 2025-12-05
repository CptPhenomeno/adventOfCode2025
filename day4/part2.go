package main

func CollectAllRollPaper(matrix Matrix) int {
	sum := 0
	for collected := CollectRollPaper(matrix); collected > 0; collected = CollectRollPaper(matrix) {
		sum += collected
	}
	return sum
}
