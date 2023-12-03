package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	read_lines := strings.Split(string(input), "\n")

	result := 0

	for lineIdx, line := range read_lines {
		lineNums := getLineNums(line)

		var prevLine string
		var nextLine string
		if lineIdx-1 >= 0 {
			prevLine = read_lines[lineIdx-1]
		}
		if lineIdx+1 <= len(read_lines)-1 {
			nextLine = read_lines[lineIdx+1]
		}

		for _, num := range lineNums {
			isPartNum := checkAroundNum(prevLine, line, nextLine, num)
			// fmt.Printf("Num: %v | %v\n", num, isPartNum)

			if isPartNum {
				numValue, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}

				result += numValue
			}
		}
	}

	fmt.Printf("Result: %v\n", result)
}

func getLineNums(line string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(line, -1)
}

func checkAroundNum(prevLine string, line string, nextLine string, num string) bool {
	numIdx := strings.Index(line, num)

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

	// Vertical check
	if prevLine != "" {
		contains := containsSymbol(prevLine[numIdx : numIdx+len(num)])
		if contains {
			return true
		}

		// Diagonal check
		if numIdx+len(num)+1 <= len(prevLine) {
			contains := containsSymbol(prevLine[numIdx : numIdx+len(num)+1])
			if contains {
				return true
			}
		}
		if numIdx-1 >= 0 {
			contains := containsSymbol(prevLine[numIdx-1 : numIdx])
			if contains {
				return true
			}
		}
	}
	if nextLine != "" {
		contains := containsSymbol(nextLine[numIdx : numIdx+len(num)])
		if contains {
			return true
		}

		// Diagonal check
		if numIdx+len(num)+1 <= len(nextLine) {
			contains := containsSymbol(nextLine[numIdx : numIdx+len(num)+1])
			if contains {
				return true
			}
		}
		if numIdx-1 >= 0 {
			contains := containsSymbol(nextLine[numIdx-1 : numIdx])
			if contains {
				return true
			}
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
