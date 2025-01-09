package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	positions := make([][]int, 0)
	velocities := make([][]int, 0)

	width, height := 101, 103

	for scanner.Scan() {
		p, v := parse(scanner.Text())
		positions = append(positions, p)
		velocities = append(velocities, v)
	}

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		render(width, height, positions)
		next(positions, velocities, width, height)
	}

}

func next(positions [][]int, velocities [][]int, width, height int) {
	for i := 0; i < len(positions); i++ {
		positions[i][0] = ((positions[i][0]+velocities[i][0])%width + width) % width
		positions[i][1] = ((positions[i][1]+velocities[i][1])%height + height) % height
	}
}

func render(width, height int, positions [][]int) {
	image := make([][]bool, height)
	for i := range image {
		image[i] = make([]bool, width)
	}
	for _, position := range positions {
		if image[position[1]][position[0]] {
			return
		}
		image[position[1]][position[0]] = true
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if image[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func parse(line string) ([]int, []int) {
	a := func(strs []string) []int {
		nums := make([]int, len(strs))
		for i, s := range strs {
			v, _ := strconv.Atoi(s)
			nums[i] = v
		}
		return nums
	}
	tmp := strings.Split(line, " ")
	return a(strings.Split(tmp[0][2:], ",")), a(strings.Split(tmp[1][2:], ","))
}
