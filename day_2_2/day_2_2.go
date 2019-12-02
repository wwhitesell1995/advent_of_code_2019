package main

import (
    "fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./day_2_2_input.txt")
	check(err)

	fileText := string(file)

	numStringList := strings.Split(fileText, ",");
	var numList []int64

	for _, num := range numStringList {
		newNum, numErr := strconv.ParseInt(num, 10, 64)
		if(numErr != nil) {
			fmt.Println("Error converting to int")
			break;
		}
		numList = append(numList, newNum)
	}

	compNum := int64(19690720)
	foundNum := false

	for i := 0; i <= 99; i++ {
		if(foundNum) {
			break
		}

		for j := 0; j<= 99; j++ {
			if(getOutput(numList, int64(i), int64(j)) == compNum) {
				fmt.Println(100 * i + j)
				foundNum = true
				break
			}
		}
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getOutput(numList []int64, noun int64, verb int64) int64 {
	currPos := 0
	currOpCode := numList[currPos]
	tempNumList := make([]int64, len(numList))
	copy(tempNumList, numList)
	tempNumList[1] = noun
	tempNumList[2] = verb

	for currOpCode != 99 {
		if(currOpCode == 1) {
			tempNumList[tempNumList[currPos+3]] = tempNumList[tempNumList[currPos+1]] + tempNumList[tempNumList[currPos+2]]
		}

		if(currOpCode == 2) {
			tempNumList[tempNumList[currPos+3]] = tempNumList[tempNumList[currPos+1]] * tempNumList[tempNumList[currPos+2]]
		}

		currPos += 4
		currOpCode = tempNumList[currPos]
	}

	return tempNumList[0]
}
