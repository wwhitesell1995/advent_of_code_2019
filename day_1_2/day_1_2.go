package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"math"
	"strconv"
)

func main() {
	totalFuel := int64(0)
	file, err := os.Open("./day_1_2_input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		parsedInt, parseErr := strconv.ParseInt(scanner.Text(), 10, 64)
		if parseErr != nil {
			log.Fatal(parseErr)
			continue;
		}
		totalFuel += getModuleResult(parsedInt)
	}
	
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
	
	fmt.Println(totalFuel)
}

func getModuleResult(value int64) int64 {
	moduleResult := int64(0)
	currFuel := getFuelResult(value)
	for currFuel > int64(0) {
	  moduleResult += currFuel
	  currFuel = getFuelResult(currFuel)
	}
	return moduleResult
}

func getFuelResult(value int64) int64 {
	fuelResult := int64(math.Floor(float64(value/3))) -2
	return fuelResult
}