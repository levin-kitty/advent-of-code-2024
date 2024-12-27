package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	antennaMap := make(map[byte][][]int)
	x := 0
	y := 0
	for scanner.Scan() {
		if y == 0 {
			y = len(scanner.Text())
		}

		for j, b := range []byte(scanner.Text()) {
			if b != '.' {
				antennaMap[b] = append(antennaMap[b], []int{x, j})
			}
		}
		x++
	}

	fmt.Println(antennaMap, x, y)

	uniquePositions := make(map[int]struct{})
	for _, antennas := range antennaMap {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				// 2a-b
				// 2b-a
				a := 2*antennas[i][0] - antennas[j][0]
				b := 2*antennas[i][1] - antennas[j][1]
				if a >= 0 && a < x && b >= 0 && b < y {
					uniquePositions[a*y+b] = struct{}{}
				}
				a = 2*antennas[j][0] - antennas[i][0]
				b = 2*antennas[j][1] - antennas[i][1]
				if a >= 0 && a < x && b >= 0 && b < y {
					uniquePositions[a*y+b] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(uniquePositions))
}
