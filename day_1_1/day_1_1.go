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
	file, err := os.Open("./day_1_1_input.txt")
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
		totalFuel += getResult(parsedInt)
	}
	
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
	
	fmt.Println(totalFuel)
}

func getResult(value int64) int64 {
	result := int64(math.Floor(float64(value/3))) -2
	return result
}