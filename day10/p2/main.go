package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	var topographicMap [][]byte
	for scanner.Scan() {
		topographicMap = append(topographicMap, []byte(scanner.Text()))
	}

	scores := 0
	for i := 0; i < len(topographicMap); i++ {
		for j := 0; j < len(topographicMap[0]); j++ {
			if topographicMap[i][j] == '0' {
				scores. += traverse(topographicMap, i, j)
			}
		}
	}
	fmt.Println(scores)
}

func traverse(topographicMap [][]byte, i, j int) int {
	if topographicMap[i][j] == '9' {
		return 1
	}

	scores := 0
	if i > 0 && topographicMap[i-1][j]-topographicMap[i][j] == 1 {
		scores += traverse(topographicMap, i-1, j)
	}
	if i < len(topographicMap)-1 && topographicMap[i+1][j]-topographicMap[i][j] == 1 {
		scores += traverse(topographicMap, i+1, j)
	}
	if j > 0 && topographicMap[i][j-1]-topographicMap[i][j] == 1 {
		scores += traverse(topographicMap, i, j-1)
	}
	if j < len(topographicMap[0])-1 && topographicMap[i][j+1]-topographicMap[i][j] == 1 {
		scores += traverse(topographicMap, i, j+1)
	}
	return scores
}
