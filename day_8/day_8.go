package main

import (
    "fmt"
	//"strings"
	"io/ioutil"
)

func main() {
	getPart1Output(25, 6)
	getPart2Output(25, 6)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getPart1Output(width int, height int) {
	file, err := ioutil.ReadFile("./day_8_input.txt")
	check(err)

	fileText := string(file)
	layerSize := width * height
	layers := make(map[int]map[string]int)
	layerNum := 0
	position := 0
	minZeroes := 99999
	minZeroLayer := 0

	for _, r := range fileText {
		if position >= layerSize {
			layerNum += 1
			position = 0
		}

		if layers[layerNum] == nil {
			layers[layerNum] = make(map[string]int)
		}

		layers[layerNum][string(r)] += 1
		position += 1
	}


	for i :=0; i <= layerNum; i++ {
		if layers[i]["0"] < minZeroes {
			minZeroes = layers[i]["0"]
			minZeroLayer = i
		}
	}

	layerOutput := layers[minZeroLayer]["1"] * layers[minZeroLayer]["2"]

	fmt.Println(layerOutput)
}

func getPart2Output(width int, height int) {
	file, err := ioutil.ReadFile("./day_8_input.txt")
	check(err)

	fileText := string(file)
	layerSize := width * height
	position := 0
	var resultLayer [150]string

	for i, _ := range resultLayer {
		resultLayer[i] = "2"
	}

	for _, r := range fileText {
		if position >= layerSize {
			position = 0
		}

		if resultLayer[position] == "2" {
			resultLayer[position] = string(r)
		}

		position += 1
	}

	// This was done to make the output a bit more readable
	for i, _ := range resultLayer {
		if resultLayer[i] == "0" {
			resultLayer[i] = " "
		}
	}

	for i:=0; i < height; i++ {
		fmt.Println(resultLayer[25*i:25*(i+1)])
	}
}