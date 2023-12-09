package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNums(line string) []int {
	splitLine := strings.Split(line, " ")
	var trimmedNums []int

	for _, num := range splitLine {
		if num == "" {
			continue
		}

		t := strings.TrimSpace(num)
		n, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}

		trimmedNums = append(trimmedNums, n)
	}

	return trimmedNums
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	read_lines := strings.Split(string(input), "\n")

	cardCounts := make(map[int]int)

	for lineIdx, line := range read_lines {
		cardValues := strings.Split(line, ": ")[1]
		valuesSplit := strings.Split(cardValues, " | ")

		winningNums := getNums(valuesSplit[0])
		myNums := getNums(valuesSplit[1])

		matchesCount := 0
		for _, num := range myNums {
			if numInSlice(num, winningNums) {
				matchesCount++
			}
		}

		cardCounts[lineIdx+1] += 1
		for k := 1; k <= cardCounts[lineIdx+1]; k++ {
			for i := 1; i <= matchesCount; i++ {
				cardCounts[lineIdx+i+1] += 1
			}
		}
		// fmt.Printf("Matches: %v %v \n", matchesCount, cardCounts)
	}

	totalCards := 0
	for _, v := range cardCounts {
		totalCards += v
	}

	fmt.Printf("Result: %v \n", totalCards)
}

func numInSlice(num int, list []int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}
