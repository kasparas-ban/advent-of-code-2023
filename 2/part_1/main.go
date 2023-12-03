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

	requiredCount := gameCount{
		red:   12,
		green: 13,
		blue:  14,
	}

	idxSum := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lineText := scanner.Text()
		lineSplit := strings.Split(lineText, ": ")
		gameIdx := strings.Split(lineSplit[0], " ")[1]
		gameIdxNum, err := strconv.Atoi(gameIdx)
		if err != nil {
			log.Fatal(err)
		}

		gameSets := strings.Split(lineSplit[1], "; ")

		meetsRequirements := true
		for _, gameSet := range gameSets {
			gameSetCounts := strings.Split(gameSet, ", ")

			counts := gameCount{
				red:   0,
				green: 0,
				blue:  0,
			}
			counts.updateCounts(gameSetCounts)

			if counts.red > requiredCount.red || counts.green > requiredCount.green || counts.blue > requiredCount.blue {
				meetsRequirements = false
			}
		}

		if meetsRequirements {
			idxSum += gameIdxNum
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v \n", idxSum)
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
