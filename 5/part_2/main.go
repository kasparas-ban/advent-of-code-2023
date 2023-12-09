package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// destination range (soil) | source range (seed) | range length

func pop(li *[][]int) []int {
	val := (*li)[len(*li)-1]
	*li = (*li)[:len(*li)-1]
	return val
}

func parseNumbers(s string) []int {
	svalues := strings.Split(s, " ")
	numbers := make([]int, len(svalues))
	for i, s := range svalues {
		numbers[i], _ = strconv.Atoi(s)
	}
	return numbers
}

func parseInput() ([][]int, [][][]int) {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	seedsLine := parseNumbers(strings.Split(lines[0], ": ")[1])
	var seeds [][]int
	for i := 0; i < len(seedsLine)-1; i += 2 {
		seeds = append(seeds, []int{seedsLine[i], seedsLine[i] + seedsLine[i+1]})
	}
	s := 0

	translators := make([][][]int, 0, 8)
	translators = append(translators, make([][]int, 0, 8))
	for i := 3; i < len(lines); i++ {
		if lines[i] == "" {
			i += 1
			s += 1
			translators = append(translators, make([][]int, 0, 8))
			continue
		}

		translators[s] = append(translators[s], parseNumbers(lines[i]))
	}

	return seeds, translators
}

func getFinalRanges(seeds *[][]int, translators [][][]int) [][]int {
	for _, t := range translators {
		var newDests [][]int

		for len(*seeds) > 0 {
			seedsRange := pop(seeds)
			sourceStart := seedsRange[0]
			sourceEnd := seedsRange[1]

			rangeAdded := false

			for _, r := range t {
				overlapStart := max(sourceStart, r[1])
				overlapEnd := min(sourceEnd, r[1]+r[2])

				if overlapStart < overlapEnd {
					newDest := []int{overlapStart - r[1] + r[0], overlapEnd - r[1] + r[0]}
					newDests = append(newDests, newDest)

					if overlapStart > sourceStart {
						startPassthrough := []int{sourceStart, overlapStart}
						*seeds = append(*seeds, startPassthrough)
					}
					if overlapEnd < sourceEnd {
						endPassthrough := []int{overlapEnd, sourceEnd}
						*seeds = append(*seeds, endPassthrough)
					}

					rangeAdded = true
					break
				}
			}

			if !rangeAdded {
				newDests = append(newDests, []int{sourceStart, sourceEnd})
			}
		}

		*seeds = newDests
	}

	return *seeds
}

func getMinLoc(ranges [][]int) int {
	minVal := 1 << 31

	for _, r := range ranges {
		if r[0] < minVal {
			minVal = r[0]
		}
	}

	return minVal
}

func main() {
	timeStart := time.Now()
	seeds, translators := parseInput()
	finalRanges := getFinalRanges(&seeds, translators)
	result := getMinLoc(finalRanges)

	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
