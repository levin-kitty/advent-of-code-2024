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

	uniquePositions := make(map[int]struct{})
	for _, antennas := range antennaMap {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				x1, y1 := antennas[i][0], antennas[i][1]
				x2, y2 := antennas[j][0], antennas[j][1]

				unit := (x2 - x1) / gcd(abs(x2-x1), abs(y2-y1))
				diff := unit * (y2 - y1) / (x2 - x1)

				fmt.Println(x1, y1, x2, y2, unit, diff)

				// start from (x1, y1) & go left
				a, b := x1, y1
				for a >= 0 && a < x && b >= 0 && b < y {
					uniquePositions[a*y+b] = struct{}{}
					a -= unit
					b -= diff
				}

				// start from (x1, y1) & go right
				a, b = x1, y1
				for a >= 0 && a < x && b >= 0 && b < y {
					uniquePositions[a*y+b] = struct{}{}
					a += unit
					b += diff
				}
			}
		}
	}

	fmt.Println(len(uniquePositions))
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
