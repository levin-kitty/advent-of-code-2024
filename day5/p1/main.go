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

	ruleMap := make(map[string]map[string]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		rule := strings.Split(line, "|")
		if _, exists := ruleMap[rule[0]]; !exists {
			ruleMap[rule[0]] = make(map[string]struct{})
		}
		ruleMap[rule[0]][rule[1]] = struct{}{}
	}

	answer := 0
	for scanner.Scan() {
		line := scanner.Text()

		pages := strings.Split(line, ",")
		isValid := true

	out:
		for i := 0; i < len(pages)-1; i++ {
			for j := i + 1; j < len(pages); j++ {
				if _, exists := ruleMap[pages[j]][pages[i]]; exists {
					isValid = false
					break out
				}
			}
		}

		if isValid {
			middle, _ := strconv.Atoi(pages[len(pages)/2])
			answer += middle
		}

	}

	fmt.Println(answer)
}
