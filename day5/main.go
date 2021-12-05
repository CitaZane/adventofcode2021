package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readInput("data.txt")
	var ventMap [1000][1000]int
	count := 0
	solveOne(input, ventMap, count)
	solveTwo(input, ventMap, count)
}
func solveTwo(input [][]int, ventMap [1000][1000]int, count int) {
	for _, vent := range input {
		x, y := vent[0], vent[1]
		for {
			ventMap[y][x]++
			if ventMap[y][x] == 2 {
				count++
			}
			if vent[2] > vent[0] {
				if x++; x > vent[2] {
					break
				}
			} else if vent[0] > vent[2] {
				if x--; x < vent[2] {
					break
				}
			}
			if vent[3] > vent[1] {
				if y++; y > vent[3] {
					break
				}
			} else if vent[1] > vent[3] {
				if y--; y < vent[3] {
					break
				}
			}
		}
	}
	fmt.Println("Result solve 2:", count)
}

//21038

func solveOne(input [][]int, ventMap [1000][1000]int, count int) {
	for _, vent := range input {
		if vent[0] != vent[2] && vent[1] != vent[3] {
			continue
		}
		x, y := vent[0], vent[1]
		for {
			ventMap[y][x]++
			if ventMap[y][x] == 2 {
				count++
			}
			if vent[2] > vent[0] {
				if x++; x > vent[2] {
					break
				}
			} else if vent[0] > vent[2] {
				if x--; x < vent[2] {
					break
				}
			}
			if vent[3] > vent[1] {
				if y++; y > vent[3] {
					break
				}
			} else if vent[1] > vent[3] {
				if y--; y < vent[3] {
					break
				}
			}
		}
	}
	fmt.Println("Result solve 1:", count)
}

// 7297

func readInput(file string) [][]int {
	f, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(f), "\n")
	var input [][]int
	for _, line := range lines {
		temp := strings.Fields(line)
		temp1 := strings.Split(temp[0], ",")
		temp2 := strings.Split(temp[2], ",")
		x1, _ := strconv.Atoi(temp1[0])
		y1, _ := strconv.Atoi(temp1[1])
		x2, _ := strconv.Atoi(temp2[0])
		y2, _ := strconv.Atoi(temp2[1])
		input = append(input, []int{x1, y1, x2, y2})
	}
	return input
}
