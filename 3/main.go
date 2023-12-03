package main

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("3/input")
	lines := bytes.Split(data, []byte("\n"))
	a(lines)
	b(lines)
}

func b(lines [][]byte) {
	regNumbers, _ := regexp.Compile("[0-9]+")
	regSymbols, _ := regexp.Compile("\\*")
	symbols := make(map[int][][]int, len(lines))
	numbers := make(map[int][][]int, len(lines))
	result := 0

	for i, line := range lines {
		symbols[i] = regSymbols.FindAllIndex(line, -1)
		numbers[i] = regNumbers.FindAllIndex(line, -1)
	}

	for i := range lines {
		for _, symbol := range symbols[i] {
			var matches []int
			//fmt.Printf("line %d symbol %v, looking for adjacent gears\n", i, symbol)
			for _, adjacent := range adjacentLineIndices(i, len(lines)-1) {
				for _, number := range numbers[adjacent] {
					if symbol[0] >= number[0]-1 && symbol[0] <= number[1] {
						matchingNumber, _ := strconv.Atoi(string(lines[adjacent][number[0]:number[1]]))
						//fmt.Printf("line %d number %v is a match\n", adjacent, number)
						matches = append(matches, matchingNumber)
					}
				}
			}
			if len(matches) == 2 {
				result += matches[0] * matches[1]
			}
		}

	}
	println(result)
}

func a(lines [][]byte) {
	regNumbers, _ := regexp.Compile("[0-9]+")
	regSymbols, _ := regexp.Compile("[^0-9.]")
	symbols := make(map[int][][]int, len(lines))
	numbers := make(map[int][][]int, len(lines))
	result := 0

	for i, line := range lines {
		symbols[i] = regSymbols.FindAllIndex(line, -1)
		numbers[i] = regNumbers.FindAllIndex(line, -1)
	}

	for i, line := range lines {
		for _, number := range numbers[i] {
			isMatch := false
			for _, adjacent := range adjacentLineIndices(i, len(lines)-1) {
				//fmt.Printf("line %d number %v check for symbols in adjacent line %d\n", i, number, adjacent)
				for _, symbol := range symbols[adjacent] {
					//fmt.Printf("line %d number %v ? symbol %v from line %d\n", i, number, symbol, adjacent)
					if symbol[0] >= number[0]-1 && symbol[0] <= number[1] {
						//fmt.Printf("line %d number %v is a match\n", i, number)
						isMatch = true
						break
					}
				}
				if isMatch {
					break
				}
			}
			if isMatch {
				matchingNumber, _ := strconv.Atoi(string(line[number[0]:number[1]]))
				result += matchingNumber
			}

		}
	}
	println(result)
}

func adjacentLineIndices(line, limit int) []int {
	low := max(0, line-1)
	high := min(limit, line+1)
	result := make([]int, high-low+1)
	for i := range result {
		result[i] = low + i
	}
	return result
}
