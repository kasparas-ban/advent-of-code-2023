package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gameCount struct {
	red   int
	green int
	blue  int
}

func main() {
	// Read file
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	result := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lineText := scanner.Text()
		lineSplit := strings.Split(lineText, ": ")
		gameSets := strings.Split(lineSplit[1], "; ")

		minGameCount := gameCount{
			red:   0,
			green: 0,
			blue:  0,
		}

		for _, gameSet := range gameSets {
			gameSetCounts := strings.Split(gameSet, ", ")

			counts := gameCount{
				red:   0,
				green: 0,
				blue:  0,
			}
			counts.updateCounts(gameSetCounts)

			if counts.red > minGameCount.red {
				minGameCount.red = counts.red
			}
			if counts.green > minGameCount.green {
				minGameCount.green = counts.green
			}
			if counts.blue > minGameCount.blue {
				minGameCount.blue = counts.blue
			}
		}

		gameResult := minGameCount.red * minGameCount.green * minGameCount.blue
		result += gameResult
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v \n", result)
}

func (g *gameCount) updateCounts(resultsLine []string) {
	for _, res := range resultsLine {
		resSplit := strings.Split(res, " ")
		count := resSplit[0]
		color := resSplit[1]

		countInt, err := strconv.Atoi(count)
		if err != nil {
			log.Fatal(err)
		}

		if color == "red" {
			g.red += countInt
		}
		if color == "green" {
			g.green += countInt
		}
		if color == "blue" {
			g.blue += countInt
		}
	}
}
