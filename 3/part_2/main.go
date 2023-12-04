package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	startIdx int
	endIdx   int
	lineIdx  int
	value    string
}

type GearLoc struct {
	x int
	y int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	read_lines := strings.Split(string(input), "\n")

	partNums := []Number{}
	gears := []GearLoc{}

	for lineIdx, line := range read_lines {
		var prevLine string
		var nextLine string
		if lineIdx-1 >= 0 {
			prevLine = read_lines[lineIdx-1]
		}
		if lineIdx+1 <= len(read_lines)-1 {
			nextLine = read_lines[lineIdx+1]
		}

		numString := ""
		numIdx := 0
		for charIdx, char := range line {
			if unicode.IsDigit(char) {
				if numString == "" {
					numIdx = charIdx
				}
				numString += string(char)

				if charIdx == len(line)-1 {
					isPartNum := checkAroundNum(prevLine, line, nextLine, numString, numIdx)
					if isPartNum {
						num := Number{
							startIdx: numIdx,
							endIdx:   numIdx + len(numString) - 1,
							lineIdx:  lineIdx,
							value:    numString,
						}
						partNums = append(partNums, num)
					}
				}
			} else {
				// Check number and add it if valid
				if numString != "" {
					isPartNum := checkAroundNum(prevLine, line, nextLine, numString, numIdx)
					if isPartNum {
						num := Number{
							startIdx: numIdx,
							endIdx:   numIdx + len(numString) - 1,
							lineIdx:  lineIdx,
							value:    numString,
						}
						partNums = append(partNums, num)
					}
				}

				numString = ""
				numIdx = charIdx
			}
		}

		// Get all gear locations
		for charIdx, char := range line {
			if char == '*' {
				gears = append(gears, GearLoc{x: charIdx, y: lineIdx})
			}
		}
	}

	var finalRes []int

	// Going over every number with gear
	for _, gear := range gears {
		numAroundGear := []Number{}
		for _, num := range partNums {
			if isGearAround(gear, num) {
				// fmt.Printf("APPENDING %v \n", numAroundGear)
				numAroundGear = append(numAroundGear, num)
			}
		}

		if len(numAroundGear) == 2 {
			val1, _ := strconv.Atoi(numAroundGear[0].value)
			val2, _ := strconv.Atoi(numAroundGear[1].value)
			finalRes = append(finalRes, val1*val2)
		}
	}

	var finalSum int
	for _, n := range finalRes {
		finalSum += n
	}

	fmt.Printf("Result: %v \n", finalSum)
}

func isGearAround(gear GearLoc, num Number) bool {
	if gear.x >= num.startIdx-1 && gear.x <= num.endIdx+1 && gear.y >= num.lineIdx-1 && gear.y <= num.lineIdx+1 {
		return true
	}
	return false
}

func checkAroundNum(prevLine string, line string, nextLine string, num string, numIdx int) bool {
	// Top check
	topStartIdx := 0
	topEndIdx := numIdx + len(num)
	if prevLine != "" {
		if numIdx-1 >= 0 {
			topStartIdx = numIdx - 1
		}
		if numIdx+len(num)+1 <= len(prevLine) {
			topEndIdx = numIdx + len(num) + 1
		}

		if containsSymbol(prevLine[topStartIdx:topEndIdx]) {
			return true
		}
	}

	// Bottom check
	bottomStartIdx := 0
	bottomEndIdx := numIdx + len(num)
	if nextLine != "" {
		if numIdx-1 >= 0 {
			bottomStartIdx = numIdx - 1
		}
		if numIdx+len(num)+1 <= len(nextLine) {
			bottomEndIdx = numIdx + len(num) + 1
		}

		if containsSymbol(nextLine[bottomStartIdx:bottomEndIdx]) {
			return true
		}
	}

	// Horizontal check
	if numIdx-1 >= 0 {
		if containsSymbol(string(line[numIdx-1])) {
			return true
		}
	}
	if numIdx+len(num) <= len(line)-1 {
		if containsSymbol(string(line[numIdx+len(num)])) {
			return true
		}
	}

	return false
}

func containsSymbol(s string) bool {
	for _, char := range s {
		if char == '*' {
			return true
		}
	}
	return false
}
