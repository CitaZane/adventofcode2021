package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, output := readInput("data.txt")
	// solveOne(output)
	solveTwo(input, output)
}

func readInput(file string) ([][]string, [][]string) {
	f, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(f), "\n")
	var input, output [][]string
	for _, line := range lines {
		temp := strings.Split(line, "|")
		inputTemp := strings.Fields(temp[0])
		outputTmep := strings.Fields(temp[1])
		input = append(input, inputTemp)
		output = append(output, outputTmep)
	}
	return input, output
}

func solveOne(output [][]string) {
	count := 0
	for _, out := range output {
		for _, o := range out {
			if len(o) == 2 || len(o) == 3 || len(o) == 4 || len(o) == 7 {
				count++
			}
		}
	}
	fmt.Println("Solve one: ", count)
}

func solveTwo(input [][]string, output [][]string) {
	var resultTot int64
	for i, example := range input {
		result := calcCase(example, output[i])
		resultTot += result
	}
	fmt.Println("Solve two: ", resultTot)

}
func decodChar(mainString string, secondString string) string {
	for _, char := range mainString {
		if !strings.Contains(secondString, string(char)) {
			return string(char)
		}
	}
	return ""
}

func findNum(input []string, base string, length int) (cleaned []string, result string) {
	for i, instance := range input {
		if len(instance) == length {
			count := 0
			for _, char := range base {
				if !strings.Contains(instance, string(char)) {
					break
				} else {
					count++
				}
			}
			if count == 4 {
				// decoder[9] = instance
				cleaned = append(input[:i], input[i+1:]...)
				return cleaned, instance
			}
		}
	}
	return cleaned, ""
}

func calcCase(input []string, output []string) int64 {
	decoder := map[int]string{0: "", 1: "", 2: "", 3: "", 4: "", 5: "", 6: "", 7: "", 8: "", 9: ""}
	segments := map[string]string{"a": "", "b": "", "c": "", "d": "", "e": "", "f": "", "g": ""}
	var cleanedInput []string
	for _, try := range input {
		if len(try) == 2 {
			decoder[1] = try
		} else if len(try) == 4 {
			decoder[4] = try
		} else if len(try) == 3 {
			decoder[7] = try
		} else if len(try) == 7 {
			decoder[8] = try
		} else {
			cleanedInput = append(cleanedInput, try)
		}
	}
	// Find A
	segments["a"] = decodChar(decoder[7], decoder[1])
	// Find 9
	temp := decoder[4] + segments["a"]
	cleanedInput, decoder[9] = findNum(cleanedInput, decoder[4], 6)
	//find g
	segments["g"] = decodChar(decoder[9], temp)
	//find e
	segments["e"] = decodChar(decoder[8], decoder[9])
	//find 3
	tempThree := decoder[1] + segments["a"] + segments["g"]
	cleanedInput, decoder[3] = findNum(cleanedInput, tempThree, 5)
	//find d
	segments["d"] = decodChar(decoder[3], tempThree)
	// find 0, 2,5,6
	for _, instance := range cleanedInput {
		if len(instance) == 6 {
			if !strings.Contains(instance, segments["d"]) {
				decoder[6] = instance
			} else {
				decoder[0] = instance
			}
		}
		if len(instance) == 5 {
			if !strings.Contains(instance, segments["e"]) {
				decoder[5] = instance
			} else {
				decoder[2] = instance
			}
		}
	}
	result := findOutputValue(decoder, output, segments)
	return result

}

func findOutputValue(decoder map[int]string, output []string, segments map[string]string) int64 {
	var tempRes []int64
	for _, digit := range output {
		switch len(digit) {
		case 2:
			tempRes = append(tempRes, 1)
		case 3:
			tempRes = append(tempRes, 7)
		case 4:
			tempRes = append(tempRes, 4)
		case 5:
			if strings.Contains(digit, segments["e"]) {
				tempRes = append(tempRes, 2)
			} else if strings.Contains(digit, string(decoder[1][0])) && strings.Contains(digit, string(decoder[1][1])) {
				tempRes = append(tempRes, 3)
			} else {
				tempRes = append(tempRes, 5)
			}
		case 6:
			if !strings.Contains(digit, segments["d"]) {
				tempRes = append(tempRes, 0)
			} else if strings.Contains(digit, string(decoder[1][0])) && strings.Contains(digit, string(decoder[1][1])) {
				tempRes = append(tempRes, 9)
			} else {
				tempRes = append(tempRes, 6)
			}
		case 7:
			tempRes = append(tempRes, 8)
		}
	}
	var result int64
	for _, count := range tempRes {
		result = result*10 + count
	}
	return result
}
