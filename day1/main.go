package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	nums := readInts("data.txt")
	firstCount := firstMesurement(nums)
	fmt.Println("First Scan: ", firstCount)
	secondCount := secondMesurement(nums)
	fmt.Println("Second Scan: ", secondCount)
}

func firstMesurement(data []int) int {
	count := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i+1] > data[i] {
			count++
		}
	}
	return count
}

func secondMesurement(data []int) int {
	count := 0
	for i := 0; i < len(data)-3; i++ {
		one := data[i]
		two := data[i+1]
		three := data[i+2]
		four := data[i+3]
		if one+two+three < two+three+four {
			count++
		}
	}
	return count
}

func readInts(file string) []int {
	d, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(d), "\n")
	nums := make([]int, 0, len(lines))
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}
	return nums
}
