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

	var tokens []string
	var rx1, ry1, rx2, ry2 int
	var ra, rb int
	var x1, y1, x2, y2 int64
	var a, b int64
	var answer int64
	for scanner.Scan() {
		tokens = strings.Split(strings.Split(scanner.Text(), ":")[1], ",")
		rx1, _ = strconv.Atoi(strings.TrimLeft(strings.Trim(tokens[0], " "), "X+"))
		ry1, _ = strconv.Atoi(strings.TrimLeft(strings.Trim(tokens[1], " "), "Y+"))

		scanner.Scan()
		tokens = strings.Split(strings.Split(scanner.Text(), ":")[1], ",")
		rx2, _ = strconv.Atoi(strings.TrimLeft(strings.Trim(tokens[0], " "), "X+"))
		ry2, _ = strconv.Atoi(strings.TrimLeft(strings.Trim(tokens[1], " "), "Y+"))

		scanner.Scan()
		tokens = strings.Split(strings.Split(scanner.Text(), ":")[1], ",")
		ra, _ = strconv.Atoi(strings.TrimLeft(strings.Trim(tokens[0], " "), "X="))
		rb, _ = strconv.Atoi(strings.TrimLeft(strings.Trim(tokens[1], " "), "Y="))

		x1, y1, x2, y2 = int64(rx1), int64(ry1), int64(rx2), int64(ry2)
		a = int64(ra) + 10000000000000
		b = int64(rb) + 10000000000000
		fmt.Println(x1, x2, a)
		fmt.Println(y1, y2, b)

		if (a*y2-b*x2)%(x1*y2-x2*y1) == 0 && (b*x1-a*y1)%(x1*y2-x2*y1) == 0 {
			x := (a*y2 - b*x2) / (x1*y2 - x2*y1)
			y := (b*x1 - a*y1) / (x1*y2 - x2*y1)
			answer += 3*x + y
			fmt.Println(x, y)
		}

		scanner.Scan()
	}

	fmt.Println(answer)
}
