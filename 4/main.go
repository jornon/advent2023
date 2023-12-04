package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, _ := os.ReadFile("4/input")
	lines := bytes.Split(data, []byte("\n"))

	b(lines)
}

func b(lines [][]byte) {
	result := int64(0)
	cardStack := make(map[int]int64, len(lines))
	cardResults := make(map[int]int, len(lines))
	spaceReg, _ := regexp.Compile(" +")
	for i, line := range lines {
		strLine := string(line)
		parts := strings.Split(strings.Split(strLine, ":")[1], "|")
		winners := spaceReg.Split(strings.TrimSpace(parts[0]), -1)
		numbers := spaceReg.Split(strings.TrimSpace(parts[1]), -1)
		cardResults[i] = countPoints(winners, numbers)
		cardStack[i] = 1
	}
	for i := 0; i < len(cardStack); i++ {
		fmt.Printf("%d x card %d has %d each, add the next %d cards to the stack\n", cardStack[i], i+1, cardResults[i], cardResults[i])
		for p := 0; p < cardResults[i]; p++ {
			if p < len(cardStack)-1 {
				fmt.Printf("add another %d x of card %d to the stack\n", cardStack[i], i+2+p)
				cardStack[i+1+p] += cardStack[i]
			}
		}
	}
	for _, v := range cardStack {
		result += v
	}
	println(result)
}

func countPoints(winners []string, numbers []string) int {
	result := 0
	for _, winner := range winners {
		if winner == " " {
			continue
		}
		for _, number := range numbers {
			if winner == number {
				result++
			}
		}
	}
	return result
}
