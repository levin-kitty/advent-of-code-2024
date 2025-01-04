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
			corners := 0
			area := 0

			for len(qi) > 0 {
				i := qi[0]
				j := qj[0]
				qi = qi[1:]
				qj = qj[1:]

				area++
				up, down, left, right := false, false, false, false

				if i > 0 && garden[i-1][j] == garden[i][j] {
					up = true
					if !isVisited[i-1][j] {
						isVisited[i-1][j] = true
						qi = append(qi, i-1)
						qj = append(qj, j)
					}
				}
				if i < n-1 && garden[i+1][j] == garden[i][j] {
					down = true
					if !isVisited[i+1][j] {
						isVisited[i+1][j] = true
						qi = append(qi, i+1)
						qj = append(qj, j)
					}
				}
				if j > 0 && garden[i][j-1] == garden[i][j] {
					left = true
					if !isVisited[i][j-1] {
						isVisited[i][j-1] = true
						qi = append(qi, i)
						qj = append(qj, j-1)
					}
				}
				if j < m-1 && garden[i][j+1] == garden[i][j] {
					right = true
					if !isVisited[i][j+1] {
						isVisited[i][j+1] = true
						qi = append(qi, i)
						qj = append(qj, j+1)
					}
				}

				// direct corner
				if !up && !left {
					corners++
				}
				if !up && !right {
					corners++
				}
				if !down && !left {
					corners++
				}
				if !down && !right {
					corners++
				}

				// indirect corner
				if (up && left) && (i > 0 && j > 0) && garden[i-1][j-1] != garden[i][j] {
					corners++
				}
				if (up && right) && (i > 0 && j < m-1) && garden[i-1][j+1] != garden[i][j] {
					corners++
				}
				if (down && left) && (i < n-1 && j > 0) && garden[i+1][j-1] != garden[i][j] {
					corners++
				}
				if (down && right) && (i < n-1 && j < m-1) && garden[i+1][j+1] != garden[i][j] {
					corners++
				}
			}

			fmt.Println(corners, area)
			answer += corners * area
		}
	}

	fmt.Println(answer)
}
