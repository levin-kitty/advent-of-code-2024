package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	var inputMap [][]byte

	x := -1
	y := -1
	direction := 0

	i := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		inputMap = append(inputMap, line)

		if x < 0 {
			for j, ch := range line {
				if ch == '^' || ch == '>' || ch == 'v' || ch == '<' {
					x = i
					y = j
				}
				switch ch {
				case '^':
					direction = 0
				case '>':
					direction = 1
				case 'v':
					direction = 2
				case '<':
					direction = 3
				}
			}
		}

		i++
	}

	fmt.Println(x, y, direction)
	visited := make([][][]bool, len(inputMap))
	for i := range visited {
		visited[i] = make([][]bool, len(inputMap[0]))
		for j := range visited[i] {
			visited[i][j] = make([]bool, len(directions))
		}
	}

	// traverse
	answer := 0
	for {
		if inputMap[x][y] != 'X' {
			inputMap[x][y] = 'X'
			answer++
		}

		if visited[x][y][direction] {
			break
		} else {
			visited[x][y][direction] = true
		}

		if x+directions[direction][0] < 0 || x+directions[direction][0] >= len(inputMap) || y+directions[direction][1] < 0 || y+directions[direction][1] >= len(inputMap[0]) {
			break
		}

		if inputMap[x+directions[direction][0]][y+directions[direction][1]] == '#' {
			direction = (direction + 1) % 4
		}

		x += directions[direction][0]
		y += directions[direction][1]
	}

	for _, line := range inputMap {
		fmt.Println(string(line))
	}

	fmt.Println(answer)
}
