package main

import (
    "fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"math"
)

func main() {
	file, err := ioutil.ReadFile("./day_3_1_input.txt")
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

func getWires(wireList []string) map[int]map[int]bool {
	currentX := 0
	currentY := 0
	wireSet := make(map[int]map[int]bool)

	for _, val := range wireList {
		runes := []rune(val)
		direction := string(runes[0])
		numIterc := strconv.ParseInt(string(runes[1:]), 10, 32)
		if numErr != nil {
			fmt.Println("Error converting to int")
			break;
		}

		for i := 0; i <= int(numIter); i++ {
			if wireSet[currentX] == nil {
				wireSet[currentX] = make(map[int]bool)
			}
			wireSet[currentX][currentY] = true
			if direction == "U" {
				currentY += 1
			} else if direction == "D" {
				currentY -= 1
			} else if direction == "R" {
				currentX += 1
			} else if direction == "L" {
				currentX -= 1
			}
		}
	}

	wireSet[0][0] = false

	return wireSet
}

func findManhattenDistance(wireSet1 map[int]map[int]bool, wireSet2 map[int]map[int]bool) int {
	minDistance := 9999999
	for key, _ := range wireSet1 {
		if wireSet2[key] == nil {
			continue
		} 

		for key1, _ := range wireSet1[key] {
			if wireSet2[key][key1] == false{
				continue
			} 

			currDistance := int(math.Abs(float64(key)) + math.Abs(float64(key1)))
			if currDistance < minDistance {
				minDistance = currDistance
			}
		}
	}
	return minDistance
}


