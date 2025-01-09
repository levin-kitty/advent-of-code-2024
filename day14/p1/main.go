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

	width, height := 101, 103
	ll, lh, hl, hh := 0, 0, 0, 0

	for scanner.Scan() {
		p, v := parse(scanner.Text())

		// after 100 seconds

		p[0] += v[0] * 100
		p[1] += v[1] * 100

		p[0] = ((p[0] % width) + width) % width
		p[1] = ((p[1] % height) + height) % height

		if p[0] < width/2 {
			if p[1] < height/2 {
				ll++
			} else if p[1] > height/2 {
				lh++
			}
		} else if p[0] > width/2 {
			if p[1] < height/2 {

				hl++
			} else if p[1] > height/2 {
				//fmt.Println(p[0], p[1])
				hh++
			}
		}
	}

	fmt.Println(ll, lh, hl, hh, ll*lh*hl*hh)

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
