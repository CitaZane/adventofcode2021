package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := readInput("data.txt")
	// solveOne(input)
	solveTwo(input)
}

func readInput(file string) []int {
	f, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(f), ",")
	var input []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		input = append(input, num)
	}
	return input
}

func solveOne(fishList []int) {
	var count int
	for i := 0; i < 18; i++ {
		for j := 0; j < len(fishList); j++ {
			switch fishList[j] {
			case 0:
				fishList[j] = 6
				count++
			default:
				fishList[j]--
			}
		}
		for l := 0; l < count; l++ {
			fishList = append(fishList, 8)
		}
		count = 0
		// fmt.Println(fishList)
	}
	fmt.Println("First solve: ", len(fishList))
}
func solveTwo(fishList []int) {
	var total uint64
	fishGrowth := map[int]uint64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	// Place first fishes
	for _, day := range fishList {
		fishGrowth[day]++
		total++
	}
	var tempCount uint64
	for i := 1; i <= 256; i++ {
		total = total + tempCount
		temp := fishGrowth[0]
		fishGrowth[0] = fishGrowth[1]
		fishGrowth[1] = fishGrowth[2]
		fishGrowth[2] = fishGrowth[3]
		fishGrowth[3] = fishGrowth[4]
		fishGrowth[4] = fishGrowth[5]
		fishGrowth[5] = fishGrowth[6]
		fishGrowth[6] = fishGrowth[7]
		fishGrowth[6] += temp
		fishGrowth[7] = fishGrowth[8]
		fishGrowth[8] = tempCount
		tempCount = fishGrowth[0]

	}
	fmt.Println("Solve two: ", total)
}
