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

func removeSpaces(s []string) []string {
	var newSlice []string
	for _, el := range s {
		if el != "" {
			newSlice = append(newSlice, el)
		}
	}

	return newSlice
}

func getRaces(input string) []race {
	lines := strings.Split(input, "\n")
	timesSplit := strings.Split(strings.Split(lines[0], ":")[1], " ")
	distancesSplit := strings.Split(strings.Split(lines[1], ":")[1], " ")

	timesSplit = removeSpaces(timesSplit)
	distancesSplit = removeSpaces(distancesSplit)

	var races []race

	for i := 0; i < len(timesSplit); i++ {
		t := strings.TrimSpace(timesSplit[i])
		d := strings.TrimSpace(distancesSplit[i])

		timeInt, err1 := strconv.Atoi(string(t))
		distInt, err2 := strconv.Atoi(string(d))
		if err1 != nil && err2 != nil {
			log.Fatal("Failed to convert to int")
		}

		race := race{time: timeInt, dist: distInt}
		races = append(races, race)
	}

	return races
}

func getWinningCount(races []race) []int {
	var winningCounts []int

	for _, race := range races {
		time, dist := race.time, race.dist

		count := 0

		// t is the button press time
		for t := 0; t <= time; t++ {
			traveledDist := (time - t) * t
			if traveledDist > dist {
				count++
			}
		}

		winningCounts = append(winningCounts, count)
	}

	return winningCounts
}

func main() {
	timeStart := time.Now()
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	races := getRaces(string(input))

	winningCounts := getWinningCount(races)

	result := winningCounts[0]
	for i := 1; i < len(winningCounts); i++ {
		result *= winningCounts[i]
	}

	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
