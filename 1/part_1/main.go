package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var numberList []string

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		var nums []int

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(line[i]) {
				digit := int(line[i] - '0')
				nums = append(nums, digit)
			}
		}

		firstNum := nums[0]
		var lastNum int
		if len(nums) > 1 {
			lastNum = nums[len(nums)-1]
		} else {
			lastNum = nums[0]
		}
		finalNum := fmt.Sprintf("%v%v", firstNum, lastNum)

		numberList = append(numberList, finalNum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for i := 0; i < len(numberList); i++ {
		num, err := strconv.Atoi(numberList[i])
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	fmt.Printf("%v \n", sum)
}
