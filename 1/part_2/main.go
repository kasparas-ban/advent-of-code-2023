package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type match struct {
	idx   int
	value int
}

var numValues = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func main() {
	var numberList []string

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var allMatches []match

		for _, key := range maps.Keys(numValues) {
			lastIdx := 0

			// Match all occurances of the given number
			for true {
				if lastIdx >= len(line) {
					break
				}

				matchedIdx := strings.Index(line[lastIdx:], key)
				fmt.Printf("\n Checking %v | key: %v | matchedIdx: %v \n", line[lastIdx+1:], key, matchedIdx)

				// Is number value matched
				if matchedIdx != -1 {
					newMatch := match{
						idx:   matchedIdx,
						value: numValues[key],
					}
					allMatches = append(allMatches, newMatch)
					lastIdx = (lastIdx + matchedIdx) + 1
				} else {
					break
				}
			}
		}

		fmt.Printf("\n Matches: %v \n", allMatches)

		// Find matched value with min idx
		firstNumber := getMin(allMatches)

		// Find matched value with max idx
		lastNumber := getMax(allMatches)

		fmt.Printf("\n FirstNum: %v | LastNum: %v \n", firstNumber, lastNumber)

		finalNum := fmt.Sprintf("%v%v", firstNumber, lastNumber)
		numberList = append(numberList, finalNum)
	}

	fmt.Printf("\n Number list %v \n", numberList)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for i := 0; i < len(numberList); i++ {
		num, err := strconv.Atoi(numberList[i])
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	fmt.Printf("\n Result: %v \n", sum)
}

func getMin(array []match) int {
	minIdx := array[0].idx
	minValue := array[0].value

	for _, element := range array {
		if element.idx < minIdx {
			minIdx = element.idx
			minValue = element.value
		}
	}

	return minValue
}

func getMax(array []match) int {
	maxIdx := array[0].idx
	maxValue := array[0].value

	for _, element := range array {
		if element.idx > maxIdx {
			maxIdx = element.idx
			maxValue = element.value
		}
	}

	return maxValue
}
