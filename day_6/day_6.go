package main

import (
    "fmt"
	"strings"
	"os"
	"bufio"
	"log"
)

func main() {
	getPart1Output()
	getPart2Output()
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func getNumOrbit(numOrbits int, pos int, currOrbit []string, orbitMap map[string][]string ) int {
	numOrbits += len(currOrbit)
	for _, val := range currOrbit {
		numOrbits = getNumOrbit(numOrbits, pos+1, orbitMap[val], orbitMap)
	}

	return numOrbits + (pos -1)
}


func findOrbitPath(pos int, currOrbit []string, orbitMap map[string][]string, searchVal string, orbitList []string ) []string {
	if contains(currOrbit, searchVal) {
		orbitList = append(orbitList, searchVal)
		return orbitList
	}

	for _, val := range currOrbit {
		orbitList = findOrbitPath(pos+1, orbitMap[val], orbitMap, searchVal, orbitList)
		if contains(orbitList, searchVal) {
			orbitList = append(orbitList, val)
			break
		}
	}

	return orbitList
}

func findOrbitDistance(orbitPath []string, divergingOrbit string, searchVal string ) int {
	divPos := 0
	searchPos := 0

	for pos, val := range orbitPath {
		if val == divergingOrbit {
			divPos = pos
		}
	}

	for pos, val := range orbitPath {
		if val == searchVal {
			searchPos = pos
		}
	}

	return searchPos - divPos - 1
}

func findDivergingOrbit(s1 []string, s2 []string) string {
	numIters := len(s1)
	prevOrbit := "COM"
	if len(s2) > len(s1) {
		numIters = len(s2)
	}

	for i := 0; i < numIters; i++ {
		if s1[i] != s2[i] {
			return prevOrbit
		}
		prevOrbit = s1[i]
	}

	return prevOrbit;
}

func getMinTransfers(pos int, currOrbit []string, orbitMap map[string][]string, orbitList []string) int {
	youOrbitPath := reverseSlice(findOrbitPath(pos, currOrbit, orbitMap, "YOU", orbitList))
	sanOrbitPath := reverseSlice(findOrbitPath(pos, currOrbit, orbitMap, "SAN", orbitList))
	divergingOrbit := findDivergingOrbit(youOrbitPath, sanOrbitPath)
	orbitDistance := findOrbitDistance(youOrbitPath, divergingOrbit, "YOU") + findOrbitDistance(sanOrbitPath, divergingOrbit, "SAN")
	return orbitDistance
}


func getPart1Output() {
	file, err := os.Open("./day_6_input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	orbitMap := make(map[string][]string)
	numOrbits := 0
	scanner := bufio.NewScanner(file)
	
    for scanner.Scan() {
		orbitString := strings.Split(scanner.Text(), ")")
		if orbitMap[orbitString[0]] == nil {
			orbitMap[orbitString[0]] = []string{orbitString[1]}
		} else {
			orbitMap[orbitString[0]] = append(orbitMap[orbitString[0]], orbitString[1])
		}
	}

	currOrbitMap := orbitMap["COM"]
	numOrbits = getNumOrbit(numOrbits, 0, currOrbitMap, orbitMap) +1

	fmt.Println(numOrbits)
}

func getPart2Output() {
	file, err := os.Open("./day_6_input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	orbitMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	
    for scanner.Scan() {
		orbitString := strings.Split(scanner.Text(), ")")
		if orbitMap[orbitString[0]] == nil {
			orbitMap[orbitString[0]] = []string{orbitString[1]}
		} else {
			orbitMap[orbitString[0]] = append(orbitMap[orbitString[0]], orbitString[1])
		}
	}

	currOrbitMap := orbitMap["COM"]
	var tempStrMap []string
	minTransfers := getMinTransfers(0, currOrbitMap, orbitMap, tempStrMap)
	fmt.Println(minTransfers)
}


