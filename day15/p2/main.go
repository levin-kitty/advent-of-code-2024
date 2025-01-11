package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	miniMap := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			doubledLine := make([]byte, 0, 2*len(line))
			for _, b := range []byte(line) {
				switch b {
				case '#':
					doubledLine = append(doubledLine, '#', '#')
				case 'O':
					doubledLine = append(doubledLine, '[', ']')
				case '.':
					doubledLine = append(doubledLine, '.', '.')
				case '@':
					doubledLine = append(doubledLine, '@', '.')
				}
			}

			miniMap = append(miniMap, []byte(doubledLine))
		}
	}
	x, y := findRobot(miniMap)
	fmt.Println(x, y)

	moves := make([]byte, 0)
	for scanner.Scan() {
		moves = append(moves, scanner.Text()...)
	}

	for _, move := range moves {
		fmt.Println("")
		fmt.Println("move:", string(move))
		x, y = tryMoveLeftRight(miniMap, x, y, move)
		//render(miniMap)
	}

	answer := 0
	for i := range miniMap {
		for j := range miniMap[i] {
			if miniMap[i][j] == '[' {
				answer += 100*i + j
			}
		}
	}
	fmt.Println(answer)
}

func findRobot(miniMap [][]byte) (int, int) {
	for i := 0; i < len(miniMap); i++ {
		for j := 0; j < len(miniMap[i]); j++ {
			if miniMap[i][j] == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func tryMoveLeftRight(miniMap [][]byte, x, y int, move byte) (int, int) {
	switch move {
	case '<':
		nextY := y - 1
		for miniMap[x][nextY] == '[' || miniMap[x][nextY] == ']' {
			nextY--
		}

		if miniMap[x][nextY] == '#' {
			return x, y
		}

		for curY := nextY; curY < y; curY++ {
			miniMap[x][curY] = miniMap[x][curY+1]
		}
		miniMap[x][y] = '.'

		return x, y - 1
	case '>':
		nextY := y + 1
		for miniMap[x][nextY] == '[' || miniMap[x][nextY] == ']' {
			nextY++
		}

		if miniMap[x][nextY] == '#' {
			return x, y
		}

		for curY := nextY; curY > y; curY-- {
			miniMap[x][curY] = miniMap[x][curY-1]
		}
		miniMap[x][y] = '.'

		return x, y + 1
	case '^':
		if tryMoveUp(miniMap, x, map[int]int{y: y}) {
			return x - 1, y
		} else {
			return x, y
		}
	case 'v':
		if tryMoveDown(miniMap, x, map[int]int{y: y}) {
			return x + 1, y
		} else {
			return x, y
		}
	}
	return -1, -1
}

func tryMoveUp(miniMap [][]byte, x int, yMap map[int]int) bool {
	nextYMap := make(map[int]int)
	allEmpty := true
	for y, _ := range yMap {
		if miniMap[x-1][y] == '#' {
			return false
		} else if miniMap[x-1][y] == '[' {
			allEmpty = false
			nextYMap[y] = y
			nextYMap[y+1] = y + 1
		} else if miniMap[x-1][y] == ']' {
			allEmpty = false
			nextYMap[y] = y
			nextYMap[y-1] = y - 1
		}
	}

	if allEmpty {
		for y, _ := range yMap {
			miniMap[x-1][y], miniMap[x][y] = miniMap[x][y], miniMap[x-1][y]
		}
		return true
	}

	if tryMoveUp(miniMap, x-1, nextYMap) {
		for y, _ := range yMap {
			miniMap[x-1][y], miniMap[x][y] = miniMap[x][y], miniMap[x-1][y]
		}
		return true
	}

	return false
}

func tryMoveDown(miniMap [][]byte, x int, yMap map[int]int) bool {
	nextYMap := make(map[int]int)
	allEmpty := true
	for y, _ := range yMap {
		if miniMap[x+1][y] == '#' {
			return false
		} else if miniMap[x+1][y] == '[' {
			allEmpty = false
			nextYMap[y] = y
			nextYMap[y+1] = y + 1
		} else if miniMap[x+1][y] == ']' {
			allEmpty = false
			nextYMap[y] = y
			nextYMap[y-1] = y - 1
		}
	}

	if allEmpty {
		for y, _ := range yMap {
			miniMap[x+1][y], miniMap[x][y] = miniMap[x][y], miniMap[x+1][y]
		}
		return true
	}

	if tryMoveDown(miniMap, x+1, nextYMap) {
		for y, _ := range yMap {
			miniMap[x+1][y], miniMap[x][y] = miniMap[x][y], miniMap[x+1][y]
		}
		return true
	}

	return false
}

func render(miniMap [][]byte) {
	for i := range miniMap {
		fmt.Println(string(miniMap[i]))
	}
}
