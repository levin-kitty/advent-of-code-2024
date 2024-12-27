package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := []byte(scanner.Text())

	diskMap := make([]int, 0, len(line))
	fileNum := 0
	for i, b := range line {
		if i%2 == 0 {
			for _ = range int(b - '0') {
				diskMap = append(diskMap, fileNum)
			}
			fileNum++
		} else {
			for _ = range int(b - '0') {
				diskMap = append(diskMap, -1)
			}
		}
	}
	fmt.Println(diskMap)

	// two pointer
	left := 0
	right := len(diskMap) - 1

	for {
		for left < len(diskMap) && diskMap[left] > -1 {
			left++
		}
		for right >= 0 && diskMap[right] == -1 {
			right--
		}
		if left >= right {
			break
		}
		diskMap[left], diskMap[right] = diskMap[right], diskMap[left]
	}

	answer := 0
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] == -1 {
			break
		}
		answer += diskMap[i] * i
	}

	fmt.Println(answer)
}
