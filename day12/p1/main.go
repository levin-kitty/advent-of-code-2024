package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	garden := make([][]byte, 0)
	for scanner.Scan() {
		garden = append(garden, []byte(scanner.Text()))
	}

	n, m := len(garden), len(garden[0])
	isVisited := make([][]bool, n)
	for i := range isVisited {
		isVisited[i] = make([]bool, m)
	}

	/* BFS */
	answer := 0
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if isVisited[x][y] {
				continue
			}

			isVisited[x][y] = true
			qi, qj := []int{x}, []int{y}
			perimeter := 0
			area := 0

			for len(qi) > 0 {
				i := qi[0]
				j := qj[0]
				qi = qi[1:]
				qj = qj[1:]

				area++
				perimeter += 4

				if i > 0 && garden[i-1][j] == garden[i][j] {
					perimeter--
					if !isVisited[i-1][j] {
						isVisited[i-1][j] = true
						qi = append(qi, i-1)
						qj = append(qj, j)
					}
				}
				if i < n-1 && garden[i+1][j] == garden[i][j] {
					perimeter--
					if !isVisited[i+1][j] {
						isVisited[i+1][j] = true
						qi = append(qi, i+1)
						qj = append(qj, j)
					}
				}
				if j > 0 && garden[i][j-1] == garden[i][j] {
					perimeter--
					if !isVisited[i][j-1] {
						isVisited[i][j-1] = true
						qi = append(qi, i)
						qj = append(qj, j-1)
					}
				}
				if j < m-1 && garden[i][j+1] == garden[i][j] {
					perimeter--
					if !isVisited[i][j+1] {
						isVisited[i][j+1] = true
						qi = append(qi, i)
						qj = append(qj, j+1)
					}
				}
			}

			answer += perimeter * area
		}
	}

	fmt.Println(answer)
}
