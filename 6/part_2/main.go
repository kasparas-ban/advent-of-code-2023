package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type race struct {
	time int
	dist int
}

func getNum(s []rune) int {
	var trimmedString []rune
	for _, el := range s {
		if el != ' ' {
			trimmedString = append(trimmedString, el)
		}
	}

	intVal, err := strconv.Atoi(string(trimmedString))
	if err != nil {
		log.Fatal(err)
	}

	return intVal
}

func getRaces(input string) race {
	lines := strings.Split(input, "\n")
	timesLine := strings.Split(lines[0], ":")[1]
	distancesLine := strings.Split(lines[1], ":")[1]

	time := getNum([]rune(timesLine))
	distance := getNum([]rune(distancesLine))

	return race{time: time, dist: distance}
}

func getWinningCount(r race) []int {
	var winningCounts []int

	time, dist := r.time, r.dist

	count := 0

	// t is the button press time
	for t := 0; t <= time; t++ {
		traveledDist := (time - t) * t
		if traveledDist > dist {
			count++
		}
	}

	winningCounts = append(winningCounts, count)

	return winningCounts
}

func main() {
	timeStart := time.Now()
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	race := getRaces(string(input))

	winningCounts := getWinningCount(race)

	result := winningCounts[0]
	for i := 1; i < len(winningCounts); i++ {
		result *= winningCounts[i]
	}

	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
