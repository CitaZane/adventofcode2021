package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Game struct {
	Board [][]int
}

type Solve struct {
	solved   bool
	game     []Game
	drawnNum []int
}

func main() {
	var solve Solve
	readInput("data.txt", &solve)
	// solveOne(0, &solve)
	solveTwo(0, &solve)
}

func readInput(file string, solve *Solve) {
	d, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(d), "\n")
	num := strings.Split(lines[0], ",")
	for _, l := range num {
		n, _ := strconv.Atoi(l)
		solve.drawnNum = append(solve.drawnNum, n)
	}

	var board Game
	for i := 2; i < len(lines); i++ {
		temp := strings.Fields(lines[i])
		var nums []int
		for _, l := range temp {
			n, _ := strconv.Atoi(l)
			nums = append(nums, n)
		}
		if len(nums) == 0 {
			continue
		}
		if len(board.Board) < 5 {
			board.Board = append(board.Board, nums)
		}
		if len(board.Board) == 5 {
			solve.game = append(solve.game, board)
			board.Board = [][]int{}
		}
	}
	return
}

func solveTwo(n int, solve *Solve) {
	if solve.solved {
		return
	}
	number := solve.drawnNum[n]
	x := len(solve.game)
	for j := 0; j < x; j++ {
		for lineIndx, line := range solve.game[j].Board {
			found, index := contains(line, number)
			if found {
				if checkRow(solve.game[j].Board, lineIndx, index) || checkColumn(solve.game[j].Board, lineIndx, index) {
					if x == 1 {
						lastNum := line[index]
						line[index] = -1
						solve.solved = true
						sum := calcSum(solve.game[0].Board)
						result := sum * lastNum
						fmt.Println(result)
						return
					}
					solve.game = append(solve.game[:j], solve.game[j+1:]...)
					j--
					x--
				} else {
					line[index] = -1
				}
			}
		}
	}
	solveTwo(n+1, solve)
}

func solveOne(n int, solve *Solve) {
	number := solve.drawnNum[n]
	if solve.solved {
		return
	}
	for _, board := range solve.game {
		for lineIndx, line := range board.Board {
			found, index := contains(line, number)
			if found {
				if checkRow(board.Board, lineIndx, index) || checkColumn(board.Board, lineIndx, index) {
					lastNum := line[index]
					line[index] = -1
					solve.solved = true
					sum := calcSum(board.Board)
					result := sum * lastNum
					fmt.Println(result)
				} else {
					line[index] = -1
				}
			}
		}

	}
	solveOne(n+1, solve)
}

func calcSum(board [][]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] != -1 {
				sum = sum + board[i][j]
			}
		}
	}
	return sum
}
func checkRow(board [][]int, line int, spot int) bool {
	for i, x := range board[line] {
		if x != -1 && i != spot {
			return false
		}
	}
	return true
}
func checkColumn(board [][]int, line int, spot int) bool {
	for j := 0; j < 5; j++ {
		if board[j][spot] != -1 && board[j][spot] != board[line][spot] {
			return false
		}
	}
	return true
}

func contains(line []int, num int) (bool, int) {
	for i := 0; i < len(line); i++ {
		if line[i] == num {
			return true, i
		}
	}
	return false, -1
}
