package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	input := readInput("data.txt")
	// solveOne(input)
	solveTwo(input)
}
func solveTwo(input []string) {
	var scoreBoard []int
	for _, line := range input {
		var temp []rune
		for index, char := range line {
			if char == '<' || char == '(' || char == '[' || char == '{' {
				temp = append(temp, char)
				if index == len(line)-1 && len(temp) != 0 {
					score := calcScore(temp)
					scoreBoard = append(scoreBoard, score)
				}
			} else if char == temp[len(temp)-1]+1 || char == temp[len(temp)-1]+2 {
				temp = temp[:len(temp)-1]
				if index == len(line)-1 && len(temp) != 0 {
					score := calcScore(temp)
					scoreBoard = append(scoreBoard, score)
				}
			} else {
				break
			}
		}
	}
	sort.Ints(scoreBoard)
	fmt.Println("Solve two:", scoreBoard[len(scoreBoard)/2])
}
func calcScore(start []rune) int {
	score := 0
	for i := len(start) - 1; i >= 0; i-- {
		char := start[i]
		score = score * 5
		if char == '(' {
			score += 1
		} else if char == '[' {
			score += 2
		} else if char == '{' {
			score += 3
		} else if char == '<' {
			score += 4
		}
	}
	return score
}

func solveOne(input []string) {
	var wrongOnes []rune
	for _, line := range input {
		var temp []rune
		for _, char := range line {
			if char == '<' || char == '(' || char == '[' || char == '{' {
				temp = append(temp, char)
			} else if char == temp[len(temp)-1]+1 || char == temp[len(temp)-1]+2 {
				temp = temp[:len(temp)-1]
			} else {
				wrongOnes = append(wrongOnes, char)
				break
			}
		}
	}
	var score int
	for _, char := range wrongOnes {
		if char == ')' {
			score += 3
		} else if char == ']' {
			score += 57
		} else if char == '}' {
			score += 1197
		} else {
			score += 25137
		}
	}
	fmt.Println("Solve one:", score)
}

func readInput(file string) []string {
	f, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(f), "\n")
	var input []string
	for _, line := range lines {
		input = append(input, line)
	}
	return input
}
