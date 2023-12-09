package main

import (
	"bufio"
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
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	totalPoints := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
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

		gamePoints := PowInts(2, matchesCount-1)

		if matchesCount == 0 {
			gamePoints = 0
		}

		totalPoints += gamePoints
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v \n", totalPoints)
}

func numInSlice(num int, list []int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}

func PowInts(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	y := PowInts(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return x * y * y
}
