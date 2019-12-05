package main

import (
    "fmt"
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./day_5_input.txt")
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

	part1Ouput := getPart1Output(numList, 1);
	part2Ouput := getPart2Output(numList, 5);
	fmt.Println(part1Ouput)
	fmt.Println(part2Ouput)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getInstructions(value int64) string {
	instruction := strconv.Itoa(int(value));
	maxIter := 5 - len(instruction)
	for i := 0; i < maxIter ; i++ {
		instruction = "0" + instruction
	}
	return instruction
}

func getPart1Output(numList []int64, inputValue int64) int64 {
	currPos := 0
	stringNum := getInstructions(numList[currPos])
	stringNumLen := len(stringNum)
	numInstructions := 4
	currOpCode, opCodeErr := strconv.ParseInt(stringNum[stringNumLen-2:], 10, 64)
	tempNumList := make([]int64, len(numList))
	output := int64(0)
	copy(tempNumList, numList)

	for currOpCode != 99 {
		numInstructions = 4
		if(currOpCode == 1) {
			inputParam := int64(currPos+3)
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]
			
			if stringNum[0:1] == "0" {
				inputParam = tempNumList[currPos+3]	
			}

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			tempNumList[inputParam] = param1 + param2
		}

		if(currOpCode == 2) {
			inputParam := int64(currPos+3)
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]
			

			if stringNum[0:1] == "0" {
				inputParam = tempNumList[currPos+3]	
			}

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			tempNumList[inputParam] = param1 * param2
		}

		if(currOpCode == 3) {
			inputParam := int64(currPos +1)
			
			if stringNum[2:3] == "0" {
				inputParam = tempNumList[currPos+1]
			}

			numInstructions = 2
			tempNumList[inputParam] = inputValue
		}

		if(currOpCode == 4) {
			inputParam := int64(currPos +1)
			
			if stringNum[2:3] == "0" {
				inputParam = tempNumList[currPos+1]
			}

			numInstructions = 2
			output = tempNumList[inputParam] 
		}

		currPos += numInstructions
		stringNum = getInstructions(tempNumList[currPos])
		numInstructions = len(strconv.Itoa(int(tempNumList[currPos])))
		stringNumLen = len(stringNum)
		currOpCode, opCodeErr = strconv.ParseInt(stringNum[stringNumLen-2:], 10, 64)
	}

	if opCodeErr != nil {
		fmt.Println("Error converting to int")
	}

	return output
}

func getPart2Output(numList []int64, inputValue int64) int64 {
	currPos := int64(0)
	stringNum := getInstructions(numList[currPos])
	stringNumLen := len(stringNum)
	currOpCode, opCodeErr := strconv.ParseInt(stringNum[stringNumLen-2:], 10, 64)
	tempNumList := make([]int64, len(numList))
	output := int64(0)
	copy(tempNumList, numList)

	for currOpCode != 99 {
		if(currOpCode == 1) {
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			tempNumList[tempNumList[currPos+3]] = param1 + param2
			currPos += int64(4)
		}

		if(currOpCode == 2) {
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			tempNumList[tempNumList[currPos+3]] = param1 * param2
			currPos += int64(4)
		}

		if(currOpCode == 3) {
			tempNumList[tempNumList[currPos+1]] = inputValue
			currPos += int64(2)
		}

		if(currOpCode == 4) {
			output = tempNumList[tempNumList[currPos+1]] 
			currPos += int64(2)
		}

		if(currOpCode == 5) {
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			if param1 != 0 {
				currPos = int64(param2)
			} else {
				currPos += 3
			}
		}

		if(currOpCode == 6) {
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			if param1 == 0 {
				currPos = int64(param2)
			} else {
				currPos += 3
			}
		}

		if(currOpCode == 7) {
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			if param1 < param2 {
				tempNumList[tempNumList[currPos+3]] = 1
			} else {
				tempNumList[tempNumList[currPos+3]] = 0
			}
			currPos += 4
		}

		if(currOpCode == 8) {
			param1 := tempNumList[currPos+1]
			param2 := tempNumList[currPos+2]

			if stringNum[2:3] == "0" {
				param1 = tempNumList[tempNumList[currPos+1]]
			}

			if stringNum[1:2] == "0" {
				param2 = tempNumList[tempNumList[currPos+2]]
			}

			if param1 == param2 {
				tempNumList[tempNumList[currPos+3]] = 1
			} else {
				tempNumList[tempNumList[currPos+3]] = 0
			}
			currPos += 4
		}

		stringNum = getInstructions(tempNumList[currPos])
		stringNumLen = len(stringNum)
		currOpCode, opCodeErr = strconv.ParseInt(stringNum[stringNumLen-2:], 10, 64)
	}

	if opCodeErr != nil {
		fmt.Println("Error converting to int")
	}

	return output
}


