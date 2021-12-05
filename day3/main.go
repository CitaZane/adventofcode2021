package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var oxygen, co2 int

func main() {
	solveOne()
	data, _ := ioutil.ReadFile("data.txt")
	dataSplit := strings.Split(string(data), "\n")
	findLifeSupport(dataSplit, 0, "oxygen")
	findLifeSupport(dataSplit, 0, "co2")
	fmt.Println(oxygen * co2)
}

func solveOne() {
	storage := make([]int, 12)
	var gamma, epsilon string
	file, _ := os.Open("data.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		for i, bit := range line {
			if bit == '0' {
				storage[i] -= 1
			} else {
				storage[i] += 1
			}
		}
	}
	for _, bit := range storage {
		if bit <= 0 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammaInt * epsilonInt)
}

func findLifeSupport(data []string, count int, rating string) {
	if len(data) == 1 {
		temporary, _ := strconv.ParseInt(data[0], 2, 64)
		if rating == "oxygen" {
			oxygen = int(temporary)
		} else {
			co2 = int(temporary)
		}
		return
	}
	common := findMostCommon(data, count, rating)
	var temp []string
	for _, number := range data {
		switch rating {
		case "oxygen":
			if number[count] == common {
				temp = append(temp, number)
			}
		case "co2":
			if number[count] != common {
				temp = append(temp, number)
			}
		}
	}
	findLifeSupport(temp, count+1, rating)
}

func findMostCommon(data []string, count int, rating string) byte {
	var sum int
	for _, line := range data {
		if line[count] == '0' {
			sum--
		} else {
			sum++
		}
	}
	if sum < 0 {
		return '0'
	} else if sum > 0 {
		return '1'
	} else {
		return '1'
	}
}
