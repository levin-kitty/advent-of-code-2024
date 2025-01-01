package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	stones := strings.Split(line, " ")

	stoneMap := make(map[string]int)
	for _, stone := range stones {
		stoneMap[stone]++
	}

	memo := make(map[string][]string)

	for i := 0; i < 75; i++ {
		fmt.Println(i)
		stoneMap = blink(stoneMap, memo)
	}

	answer := 0
	for _, v := range stoneMap {
		answer += v
	}

	fmt.Println(answer)
}

func blink(oldStoneMap map[string]int, memo map[string][]string) map[string]int {
	newStoneMap := make(map[string]int)

	for stone, freq := range oldStoneMap {
		if result, ok := memo[stone]; ok {
			for _, r := range result {
				newStoneMap[r] += freq
			}
			continue
		}

		var result []string
		if stone == "0" {
			result = append(result, "1")
		} else if len(stone)%2 == 0 {
			half := len(stone) / 2
			result = append(result, stone[:half])

			if number, _ := strconv.Atoi(stone[half:]); number == 0 {
				result = append(result, "0")
			} else {
				result = append(result, strings.TrimLeft(stone[half:], "0"))
			}
		} else {
			number, _ := strconv.Atoi(stone)
			result = append(result, strconv.Itoa(number*2024))
		}

		memo[stone] = result
		for _, r := range result {
			newStoneMap[r] += freq
		}
	}

	return newStoneMap
}
