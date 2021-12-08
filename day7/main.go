package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Stats struct {
	largest   int
	steps     map[int]int
	positions []int
}

var stats Stats

func main() {
	// readInts("test.txt")
	readInts("data.txt")
	sort.Ints(stats.positions)
	stats.largest = stats.positions[len(stats.positions)-1]
	stats.steps = make(map[int]int)
	// solveOne()
	solveTwo()
}

func solveOne() {
	size := stats.largest - stats.positions[0]
	for i := stats.positions[0]; i < size; i++ {
		count := 0
		for _, value := range stats.positions {
			if value > i {
				count += value - i
			} else {
				count += i - value
			}
		}
		stats.steps[i] = count
	}
	max := 1000000
	res := -1
	for key, value := range stats.steps {
		if value < max {
			res = key
			max = value
		}
	}
	fmt.Println(stats.steps[res])
}

func solveTwo() {
	target := (stats.largest - stats.positions[0]) / 2
	fmt.Println(target)
	var bestSoFar int
	var bestResult int
	var direction string
	// stats.steps[target]
	bestResult = calculateCost(target)
	bestSoFar = target
	resPlus := calculateCost(target + 1)
	resMinus := calculateCost(target - 1)
	if resPlus < bestResult {
		bestSoFar = target + 1
		bestResult = resPlus
		direction = "+"
	}
	if resMinus < bestResult {
		bestSoFar = target - 1
		bestResult = resMinus
		direction = "-"
	}
	switch direction {
	case "-":
		for {
			next := calculateCost(bestSoFar - 1)
			if next > bestResult {
				break
			} else {
				bestResult = next
				bestSoFar -= 1
				next = calculateCost(bestSoFar - 1)
			}
		}
	case "+":
		for {
			next := calculateCost(bestSoFar + 1)
			if next > bestResult {
				break
			} else {
				bestResult = next
				bestSoFar += 1
				next = calculateCost(bestSoFar + 1)
			}
		}
	}
	fmt.Println(bestSoFar, bestResult, direction)
}
func calculateCost(target int) (res int) {
	for _, value := range stats.positions {
		currCost := 0
		if value > target {
			i := 0
			for {
				if value == target {
					break
				}
				value--
				i++
				currCost += i
			}
		} else if value < target {
			i := 0
			for {
				if value == target {
					break
				}
				value++
				i++
				currCost += i
			}
		}
		res += currCost
	}
	return res
}

func readInts(file string) {
	d, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(d), ",")
	stats.positions = make([]int, 0, len(lines))
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, _ := strconv.Atoi(l)
		stats.positions = append(stats.positions, n)
	}
}
