package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// destination range (soil) | source range (seed) | range length

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// secRes := make(map[int]int)

	// Get seeds
	all_lines := strings.Split(string(input), "\n")
	seeds := toIntSlice(strings.Split(strings.Split(all_lines[0], ": ")[1], " "))

	// Get almanac values
	sections := strings.Split(string(input), "\n\n")[1:]

	lowestLoc := math.MaxInt

	for _, seed := range seeds {
		destValue := getMapValue(seed, 0, sections[0])
		for secIdx := 1; secIdx < len(sections); secIdx++ {
			destValue = getMapValue(destValue, secIdx, sections[secIdx])
		}

		if destValue < lowestLoc {
			lowestLoc = destValue
		}
	}

	fmt.Printf("Result: %v \n", lowestLoc)
}

func toIntSlice(slice []string) []int {
	intSlice := []int{}
	for _, s := range slice {
		intS, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		intSlice = append(intSlice, intS)
	}

	return intSlice
}

func getMapValue(source int, secIdx int, sec string) int {
	lines := strings.Split(sec, "\n")[1:]

	destVal := source

	for _, line := range lines {
		lineVals := toIntSlice(strings.Split(line, " "))
		dest := lineVals[0]
		s := lineVals[1]
		r := lineVals[2]

		if source >= s && source < s+r {
			destVal = dest + (source - s)
		}
	}

	return destVal
}
