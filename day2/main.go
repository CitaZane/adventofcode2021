package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstSolve()
	secondSolve()
}

func firstSolve() {
	var x, y int
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, " ")
		switch v[0] {
		case "forward":
			value, _ := strconv.Atoi(v[1])
			x = x + value
		case "down":
			value, _ := strconv.Atoi(v[1])
			y = y + value
		case "up":
			value, _ := strconv.Atoi(v[1])
			y = y - value
		default:
			fmt.Println(x, y)
		}
	}
	file.Close()
	fmt.Println(x * y)
}
func secondSolve() {
	var x, y, a int
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, " ")
		switch v[0] {
		case "forward":
			value, _ := strconv.Atoi(v[1])
			x = x + value
			y = y + value*a
		case "down":
			value, _ := strconv.Atoi(v[1])
			a = a + value
		case "up":
			value, _ := strconv.Atoi(v[1])
			a = a - value
		default:
			fmt.Println(x, y, a)
		}
	}
	file.Close()
	fmt.Println(x * y)
}
