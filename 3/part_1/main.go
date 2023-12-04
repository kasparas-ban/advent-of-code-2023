package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	read_lines := strings.Split(string(input), "\n")

	result := 0

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
						numValue, err := strconv.Atoi(numString)
						if err != nil {
							log.Fatal(err)
						}

						result += numValue
					}
				}
			} else {
				// Check number and add it if valid
				if numString != "" {
					isPartNum := checkAroundNum(prevLine, line, nextLine, numString, numIdx)
					if isPartNum {
						numValue, err := strconv.Atoi(numString)
						if err != nil {
							log.Fatal(err)
						}

						result += numValue
					}
				}

				numString = ""
				numIdx = charIdx
			}
		}
	}

	fmt.Printf("Result: %v\n", result)
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
		if !unicode.IsDigit(char) && char != '.' {
			return true
		}
	}
	return false
}
