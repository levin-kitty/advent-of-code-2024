package day4

import (
	"bufio"
	"os"
)

func day4Input() []string {
	file, err := os.Open("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func day4Part1() int {
	lines := day4Input()
	answer := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if j+3 < len(lines) && lines[i][j] == 'X' && lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
				answer++
			}
			if j+3 < len(lines) && lines[i][j] == 'S' && lines[i][j+1] == 'A' && lines[i][j+2] == 'M' && lines[i][j+3] == 'X' {
				answer++
			}
			if i+3 < len(lines) && lines[i][j] == 'X' && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
				answer++
			}
			if i+3 < len(lines) && lines[i][j] == 'S' && lines[i+1][j] == 'A' && lines[i+2][j] == 'M' && lines[i+3][j] == 'X' {
				answer++
			}
			if j+3 < len(lines) && i+3 < len(lines) && lines[i][j] == 'X' && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
				answer++
			}
			if j+3 < len(lines) && i+3 < len(lines) && lines[i][j] == 'S' && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'M' && lines[i+3][j+3] == 'X' {
				answer++
			}
			if i+3 < len(lines) && j-3 >= 0 && lines[i][j] == 'X' && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
				answer++
			}
			if i+3 < len(lines) && j-3 >= 0 && lines[i][j] == 'S' && lines[i+1][j-1] == 'A' && lines[i+2][j-2] == 'M' && lines[i+3][j-3] == 'X' {
				answer++
			}
		}
	}

	return answer
}

func day4Part2() int {
	lines := day4Input()
	answer := 0

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' {
				if ((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')) && ((lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S') || (lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M')) {
					answer++
				}
			}
		}
	}

	return answer
}
