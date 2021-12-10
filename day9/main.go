package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var ventMap [][]int

type Basin struct {
	size           int
	lowPoint       []int
	waitingRoom    [][]int
	alreadyVisited [][]int
}

func main() {
	// ventMap = readInts("test.txt")
	ventMap = readInts("data.txt")
	// fmt.Println(ventMap)
	coordinates := solveOne()
	solveTwo(coordinates)
}
func solveOne() [][]int {
	lowPoints := 0
	count := 0
	var lowpointCoordinates [][]int
	edge := len(ventMap[0]) - 1
	for lineIdx, line := range ventMap {
		for pointIdx, point := range line {
			switch lineIdx {
			case 0:
				if pointIdx == edge {
					if point < line[pointIdx-1] && point < ventMap[lineIdx+1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				} else if pointIdx == 0 {
					if point < line[pointIdx+1] && point < ventMap[lineIdx+1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				} else {
					if point < line[pointIdx+1] && point < line[pointIdx-1] && point < ventMap[lineIdx+1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				}
			case len(ventMap) - 1:
				if pointIdx == edge {
					if point < line[pointIdx-1] && point < ventMap[lineIdx-1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				} else if pointIdx == 0 {
					if point < line[pointIdx+1] && point < ventMap[lineIdx-1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				} else {
					if point < line[pointIdx+1] && point < line[pointIdx-1] && point < ventMap[lineIdx-1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				}
			default:
				if pointIdx == edge {
					if point < line[pointIdx-1] && point < ventMap[lineIdx-1][pointIdx] && point < ventMap[lineIdx+1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				} else if pointIdx == 0 {
					if point < line[pointIdx+1] && point < ventMap[lineIdx-1][pointIdx] && point < ventMap[lineIdx+1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				} else {
					if point < line[pointIdx+1] && point < line[pointIdx-1] && point < ventMap[lineIdx-1][pointIdx] && point < ventMap[lineIdx+1][pointIdx] {
						lowpointCoordinates = append(lowpointCoordinates, []int{lineIdx, pointIdx})
						lowPoints += point
						count++
					}
				}
			}
		}

	}
	fmt.Println("Solve One: ", lowPoints+count)
	return lowpointCoordinates
}
func solveTwo(lowPoint [][]int) {
	var basins []int
	for index, point := range lowPoint {
		var basin Basin
		basin.lowPoint = lowPoint[index]
		basin.waitingRoom = append(basin.waitingRoom, point)
		findBasinSize(&basin)
		basins = append(basins, basin.size)
	}
	findResult(basins)

}
func findBasinSize(basin *Basin) {
	if len(basin.waitingRoom) == 0 {
		return
	}
	active := basin.waitingRoom[0]
	basin.waitingRoom = basin.waitingRoom[1:]
	basin.size++
	basin.alreadyVisited = append(basin.alreadyVisited, active)
	// Append all 4 sides to waiting room
	if active[0] > 0 {
		up := []int{active[0] - 1, active[1]}
		if ventMap[up[0]][up[1]] != 9 && notVisited(up, basin.alreadyVisited) && notWaiting(up, basin.waitingRoom) {
			basin.waitingRoom = append(basin.waitingRoom, up)
		}
	}
	if active[1] < len(ventMap[0])-1 {
		right := []int{active[0], active[1] + 1}
		if ventMap[right[0]][right[1]] != 9 && notVisited(right, basin.alreadyVisited) && notWaiting(right, basin.waitingRoom) {
			basin.waitingRoom = append(basin.waitingRoom, right)
		}

	}
	if active[0] < len(ventMap)-1 {
		down := []int{active[0] + 1, active[1]}
		if ventMap[down[0]][down[1]] != 9 && notVisited(down, basin.alreadyVisited) && notWaiting(down, basin.waitingRoom) {
			basin.waitingRoom = append(basin.waitingRoom, down)
		}

	}
	if active[1] > 0 {
		left := []int{active[0], active[1] - 1}
		if ventMap[left[0]][left[1]] != 9 && notVisited(left, basin.alreadyVisited) && notWaiting(left, basin.waitingRoom) {
			basin.waitingRoom = append(basin.waitingRoom, left)
		}

	}
	findBasinSize(basin)
}
func notVisited(coordinate []int, list [][]int) bool {
	for _, point := range list {
		if reflect.DeepEqual(coordinate, point) {
			return false
		}
	}
	return true
}
func notWaiting(coordinate []int, list [][]int) bool {
	for _, point := range list {
		if reflect.DeepEqual(coordinate, point) {
			return false
		}
	}
	return true
}
func findResult(basins []int) {
	sort.Ints(basins)
	res := basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
	fmt.Println("Solve two: ", res)

}
func readInts(file string) [][]int {
	f, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(f), "\n")
	nums := make([][]int, 0, len(lines))
	for _, line := range lines {
		var temp []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			temp = append(temp, num)
		}
		nums = append(nums, temp)
	}
	return nums
}
