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
				memo := make(map[int]struct{})
				traverse(topographicMap, i, j, memo)
				scores += len(memo)
			}
		}
	}
	fmt.Println(scores)
}

func traverse(topographicMap [][]byte, i, j int, memo map[int]struct{}) {
	if topographicMap[i][j] == '9' {
		memo[i*len(topographicMap)+j] = struct{}{}
	}

	if i > 0 && topographicMap[i-1][j]-topographicMap[i][j] == 1 {
		traverse(topographicMap, i-1, j, memo)
	}
	if i < len(topographicMap)-1 && topographicMap[i+1][j]-topographicMap[i][j] == 1 {
		traverse(topographicMap, i+1, j, memo)
	}
	if j > 0 && topographicMap[i][j-1]-topographicMap[i][j] == 1 {
		traverse(topographicMap, i, j-1, memo)
	}
	if j < len(topographicMap[0])-1 && topographicMap[i][j+1]-topographicMap[i][j] == 1 {
		traverse(topographicMap, i, j+1, memo)
	}
}
