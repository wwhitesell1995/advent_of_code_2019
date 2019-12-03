package main

import (
    "fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./day_3_2_input.txt")
	check(err)

	fileText := string(file)

	wireList := strings.Split(fileText, "\n");

	wire1Directions := strings.Split(wireList[0], ",")
	wire2Directions := strings.Split(wireList[1], ",")

	wireSet1 := getWires(wire1Directions)
	wireSet2 := getWires(wire2Directions)
	minDistance := findManhattenDistance(wireSet1, wireSet2)

	fmt.Println(minDistance)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getWires(wireList []string) map[int]map[int]int {
	currentX := 0
	currentY := 0
	numSteps := 0
	wireSet := make(map[int]map[int]int)

	for _, val := range wireList {
		runes := []rune(val)
		direction := string(runes[0])
		numIter, numErr := strconv.ParseInt(string(runes[1:]), 10, 32)
		if numErr != nil {
			fmt.Println("Error converting to int")
			break;
		}

		for i := 0; i < int(numIter); i++ {
			if wireSet[currentX] == nil {
				wireSet[currentX] = make(map[int]int)
			}

			if wireSet[currentX][currentY] == 0 || wireSet[currentX][currentY] > numSteps {
				wireSet[currentX][currentY] = numSteps
			}
			
			if direction == "U" {
				currentY += 1
			} else if direction == "D" {
				currentY -= 1
			} else if direction == "R" {
				currentX += 1
			} else if direction == "L" {
				currentX -= 1
			}
			numSteps += 1
		}
	}

	wireSet[0][0] = 0

	return wireSet
}

func findManhattenDistance(wireSet1 map[int]map[int]int, wireSet2 map[int]map[int]int) int {
	minSteps := 9999999
	for key, _ := range wireSet1 {
		if wireSet2[key] == nil {
			continue
		} 

		for key1, _ := range wireSet1[key] {
			if wireSet2[key][key1] == 0 {
				continue
			} 

			stepTotal := wireSet1[key][key1] + wireSet2[key][key1]
			if stepTotal < minSteps {
				minSteps = stepTotal
			}
		}
	}
	return minSteps
}


