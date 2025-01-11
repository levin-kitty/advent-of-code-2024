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
			miniMap = append(miniMap, []byte(line))
		}
	}
	x, y := findRobot(miniMap)

	moves := make([]byte, 0)
	for scanner.Scan() {
		moves = append(moves, scanner.Text()...)
	}

	render(miniMap)
	for _, move := range moves {
		x, y = tryMove(miniMap, x, y, move)
	}
	render(miniMap)

	answer := 0
	for i := range miniMap {
		for j := range miniMap[i] {
			if miniMap[i][j] == 'O' {
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

func tryMove(miniMap [][]byte, x, y int, move byte) (int, int) {
	switch move {
	case '<':
		nextY := y - 1
		for miniMap[x][nextY] == 'O' {
			nextY--
		}

		if miniMap[x][nextY] == '#' {
			return x, y
		}

		miniMap[x][y] = '.'
		miniMap[x][y-1] = '@'
		for curY := nextY; curY < y-1; curY++ {
			miniMap[x][curY] = 'O'
		}

		return x, y - 1
	case '>':
		nextY := y + 1
		for miniMap[x][nextY] == 'O' {
			nextY++
		}

		if miniMap[x][nextY] == '#' {
			return x, y
		}

		miniMap[x][y] = '.'
		miniMap[x][y+1] = '@'
		for curY := nextY; curY > y+1; curY-- {
			miniMap[x][curY] = 'O'
		}
		return x, y + 1
	case '^':
		nextX := x - 1
		for miniMap[nextX][y] == 'O' {
			nextX--
		}

		if miniMap[nextX][y] == '#' {
			return x, y
		}

		miniMap[x][y] = '.'
		miniMap[x-1][y] = '@'
		for curX := nextX; curX < x-1; curX++ {
			miniMap[curX][y] = 'O'
		}

		return x - 1, y
	case 'v':
		nextX := x + 1
		for miniMap[nextX][y] == 'O' {
			nextX++
		}

		if miniMap[nextX][y] == '#' {
			return x, y
		}

		miniMap[x][y] = '.'
		miniMap[x+1][y] = '@'
		for curX := nextX; curX > x+1; curX-- {
			miniMap[curX][y] = 'O'
		}
		return x + 1, y
	}
	return -1, -1
}

func render(miniMap [][]byte) {
	for i := range miniMap {
		fmt.Println(string(miniMap[i]))
	}
}
