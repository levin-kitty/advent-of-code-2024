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

	xOrigin := -1
	yOrigin := -1
	directionOrigin := 0

	i := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		inputMap = append(inputMap, line)

		if xOrigin < 0 {
			for j, ch := range line {
				if ch == '^' || ch == '>' || ch == 'v' || ch == '<' {
					xOrigin = i
					yOrigin = j
				}
				switch ch {
				case '^':
					directionOrigin = 0
				case '>':
					directionOrigin = 1
				case 'v':
					directionOrigin = 2
				case '<':
					directionOrigin = 3
				}
			}
		}

		i++
	}

	// traverse
	answer := 0
	for i := 0; i < len(inputMap); i++ {
		for j := 0; j < len(inputMap[0]); j++ {

			if inputMap[i][j] == '#' || (i == xOrigin && j == yOrigin) {
				continue
			}
			copiedMap := deepCopy(inputMap)
			copiedMap[i][j] = '#'

			x := xOrigin
			y := yOrigin
			direction := directionOrigin

			visited := make([][][]bool, len(copiedMap))
			for i := range visited {
				visited[i] = make([][]bool, len(copiedMap[0]))
				for j := range visited[i] {
					visited[i][j] = make([]bool, len(directions))
				}
			}

			var isLoop bool

			for {
				if copiedMap[x][y] != 'X' {
					copiedMap[x][y] = 'X'
				}

				if visited[x][y][direction] {
					isLoop = true
					break
				} else {
					visited[x][y][direction] = true
				}

				if x+directions[direction][0] < 0 || x+directions[direction][0] >= len(inputMap) || y+directions[direction][1] < 0 || y+directions[direction][1] >= len(inputMap[0]) {
					isLoop = false
					break
				}

				if copiedMap[x+directions[direction][0]][y+directions[direction][1]] == '#' {
					direction = (direction + 1) % 4
				} else {
					x += directions[direction][0]
					y += directions[direction][1]
				}
			}

			if isLoop {
				fmt.Println(i, j)
				answer++
			}
		}
	}

	fmt.Println(answer)
}

func deepCopy(origin [][]byte) [][]byte {
	copied := make([][]byte, len(origin))
	for i := range copied {
		copied[i] = make([]byte, len(origin[i]))
		copy(copied[i], origin[i])
	}
	return copied
}
