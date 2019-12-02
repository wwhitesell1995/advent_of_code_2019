package main

import (
    "fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./day_2_1_input.txt")
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

	currPos := 0
	currOpCode := numList[currPos]

	for currOpCode != 99 {
		if(currOpCode == 1) {
			numList[numList[currPos+3]] = numList[numList[currPos+1]] + numList[numList[currPos+2]]
		}

		if(currOpCode == 2) {
			numList[numList[currPos+3]] = numList[numList[currPos+1]] * numList[numList[currPos+2]]
		}

		currPos += 4
		currOpCode = numList[currPos]
	}

	fmt.Println(numList[0])
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}