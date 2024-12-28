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

	fileIndex := make([]int, 0)
	fileSizes := make([]int, 0)
	emptyIndex := make([]int, 0)
	emptySizes := make([]int, 0)

	diskMap := make([]int, 0, len(line))
	fileNum := 0
	for i, b := range line {
		if i%2 == 0 {
			fileIndex = append(fileIndex, len(diskMap))
			for _ = range int(b - '0') {
				diskMap = append(diskMap, fileNum)
			}
			fileSizes = append(fileSizes, int(b-'0'))
			fileNum++
		} else {
			emptyIndex = append(emptyIndex, len(diskMap))
			emptySizes = append(emptySizes, int(b-'0'))
			for _ = range int(b - '0') {
				diskMap = append(diskMap, -1)
			}
		}
	}

	fmt.Println(diskMap)

	for f := fileNum - 1; f >= 0; f-- {
		for e := 0; e < len(emptyIndex); e++ {
			if emptyIndex[e] >= fileIndex[f] {
				break
			}

			if emptySizes[e] >= fileSizes[f] {
				for i := 0; i < fileSizes[f]; i++ {
					diskMap[i+emptyIndex[e]] = f
				}
				for i := 0; i < fileSizes[f]; i++ {
					diskMap[i+fileIndex[f]] = -1
				}

				emptySizes[e] -= fileSizes[f]
				emptyIndex[e] += fileSizes[f]
				break
			}
		}
	}

	answer := 0
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] > -1 {
			answer += diskMap[i] * i
		}
	}

	fmt.Println(answer)
}
