package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/maps"
)

type game struct {
	hand string
	bid  int
}

var CardValueMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func getHandCategory(hand string) int {
	cardCounts := make(map[rune]int)
	for _, c := range hand {
		cardCounts[c]++
	}
	maxCount := 0
	for _, value := range cardCounts {
		if value > maxCount {
			maxCount = value
		}
	}
	cardValues := maps.Values(cardCounts)

	// 1. Five of a kind
	if maxCount == 5 {
		return 1
	}

	// 2. Four of a kind
	if maxCount == 4 {
		return 2
	}

	// 3. Full house
	if maxCount == 3 && len(cardValues) == 2 {
		return 3
	}

	// 4. Three of a kind
	if len(maps.Values(cardCounts)) == 3 {
		if maxCount == 3 {
			return 4
		}
		// 5. Two pair
		if maxCount == 2 {
			return 5
		}
	}

	if maxCount == 1 {
		return 7
	}

	return 6
}

func getGames(input string) []game {
	lines := strings.Split(input, "\n")

	var games []game
	for _, line := range lines {
		res := strings.Split(line, " ")
		bidVal, err := strconv.Atoi(res[1])
		if err != nil {
			log.Fatal(err)
		}
		games = append(games, game{hand: res[0], bid: bidVal})
	}

	return games
}

func sortGames(games []game) []game {
	sortedGames := games
	sort.Slice(games, func(i, j int) bool {
		for k := 0; k < 5; k++ {
			char1 := rune(games[i].hand[k])
			char2 := rune(games[j].hand[k])
			if char1 == char2 {
				continue
			}
			return CardValueMap[char1] > CardValueMap[char2]
		}

		return true
	})

	return sortedGames
}

func main() {
	timeStart := time.Now()
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games := getGames(string(input))

	categoriesMap := make(map[int][]game)
	for _, g := range games {
		cat := getHandCategory(g.hand)
		categoriesMap[cat] = append(categoriesMap[cat], g)
	}

	categoriesSorted := maps.Keys(categoriesMap)
	slices.Sort(categoriesSorted)

	var allSortedGames []game
	for _, c := range categoriesSorted {
		sortedGames := sortGames(categoriesMap[c])
		allSortedGames = append(allSortedGames, sortedGames...)
	}

	result := 0
	for i := 1; i < len(allSortedGames)+1; i++ {
		result += (len(allSortedGames) + 1 - i) * allSortedGames[i-1].bid
	}

	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
