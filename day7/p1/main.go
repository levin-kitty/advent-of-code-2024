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

	atoi := func(strs []string) []int64 {
		ints := make([]int64, len(strs))
		for i, str := range strs {
			t, _ := strconv.Atoi(str)
			ints[i] = int64(t)
		}
		return ints
	}

	answer := int64(0)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ":")

		t, _ := strconv.Atoi(tokens[0])
		target := int64(t)

		values := atoi(strings.Split(tokens[1], " ")[1:])

		if isValidEquation(values, 1, values[0], target) {
			fmt.Println(target, values)
			answer += target
		}
	}

	fmt.Println("answer:", answer)
}

func isValidEquation(values []int64, current int, currentValue int64, target int64) bool {

	/* terminal conditions */

	if current == len(values) {
		return currentValue == target
	} else if currentValue > target {
		return false
	}

	/* recursion */

	if isValidEquation(values, current+1, currentValue+values[current], target) {
		return true
	}
	if isValidEquation(values, current+1, currentValue*values[current], target) {
		return true
	}
	return false
}
