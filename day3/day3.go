package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func day3Input() []string {
	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func day3Part1() int {
	reExpr := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	reNum := regexp.MustCompile(`[0-9]{1,3}`)
	solve := func(line string) (answer int) {
		validExprs := reExpr.FindAllString(line, -1)
		for _, expr := range validExprs {
			nums := reNum.FindAllString(expr, 2)
			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])
			answer += n1 * n2
		}
		return
	}

	answer := 0
	lines := day3Input()
	for _, line := range lines {
		answer += solve(line)
	}

	return answer
}

func day3Part2() int {
	reExpr := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	//reExpr := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|(d)(o)(?:n't)?\(\)`)
	reNum := regexp.MustCompile(`\d{1,3}`)
	solve := func(line string) (answer int) {
		validExprs := reExpr.FindAllString(line, -1)
		fmt.Println(validExprs)
		do := true
		for _, expr := range validExprs {
			switch expr {
			case `do()`:
				do = true
			case `don't()`:
				do = false
			default:
				if do {
					nums := reNum.FindAllString(expr, 2)
					n1, _ := strconv.Atoi(nums[0])
					n2, _ := strconv.Atoi(nums[1])
					answer += n1 * n2
				} else {
					fmt.Println("don't(): ", expr)
				}
			}
		}
		return
	}

	return solve(strings.Join(day3Input(), ""))
}
