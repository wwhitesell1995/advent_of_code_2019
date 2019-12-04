package main

import (
    "fmt"
	"strconv"
)

func main() {

	part1Solution := part1solution()
	part2Solution := part2solution()

	fmt.Println(part1Solution)
	fmt.Println(part2Solution)
}

func part1solution() int {
	minValue := 138307
	maxValue := 654504
	numPasswords := 0

	for i := minValue; i <= maxValue; i++ {
		stringNum := strconv.Itoa(i)
		isAdjacent := false
		isGreater := true
		maxNum := 0
		prevNum := 0
		for _, char := range stringNum {
			digit := int(char - '0')
			if digit >= maxNum {
				maxNum = digit
			} else {
				isGreater = false
			}

			if digit == prevNum {
				isAdjacent = true
			}

			prevNum = digit
		}

		if isAdjacent && isGreater {
			numPasswords += 1
		}
	}

	return numPasswords
}

func part2solution() int {
	minValue := 138307
	maxValue := 654504
	numPasswords := 0

	for i := minValue; i <= maxValue; i++ {
		stringNum := strconv.Itoa(i)
		isNumMatched := false
		isGreater := true
		maxNum := 0
		prevNum := 0
		numMatched := 0
		for pos, char := range stringNum {
			digit := int(char - '0')
			if digit >= maxNum {
				maxNum = digit
			} else {
				isGreater = false
			}

			if digit == prevNum {
				numMatched += 1
			}

			if digit != prevNum {
				if numMatched == 1 {
					isNumMatched = true
				}
				numMatched = 0
			}

			if numMatched == 1 && pos == len(stringNum) - 1 {
				isNumMatched = true
			}

			prevNum = digit
		}

		if isNumMatched && isGreater {
			numPasswords += 1
		}
	}

	return numPasswords
}


