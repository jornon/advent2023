package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("input")
	lines := bytes.Split(data, []byte("\n"))
	a(lines)
	b(lines)
}

func a(lines [][]byte) {
	reg, _ := regexp.Compile("[0-9]")
	result := 0
	for _, line := range lines {
		matches := reg.FindAll(line, -1)
		cal, _ := strconv.Atoi(fmt.Sprintf("%s%s", matches[0], matches[len(matches)-1]))
		result += cal
	}
	fmt.Printf("Result A: %d\n", result)
}

func b(lines [][]byte) {
	result := 0
	for _, line := range lines {
		matches := interpreter(line)
		first := toNum(string(matches[0]))
		last := toNum(string(matches[len(matches)-1]))
		cal, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		result += cal
		_, _ = fmt.Printf("%s: %s %s = %d (%d)\n", string(line), first, last, cal, result)
	}
	fmt.Printf("Result B: %d\n", result)
}

func toNum(s string) string {
	if numMap[s] == "" {
		return s
	} else {
		return numMap[s]
	}
}

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func interpreter(lines []byte) (matches [][]byte) {
	reg, _ := regexp.Compile("[0-9]|one|two|three|four|five|six|seven|eight|nine|ten")
	loc := reg.FindIndex(lines)
	if loc != nil {
		matches = append(matches, lines[loc[0]:loc[1]])
		matches = append(matches, interpreter(lines[loc[0]+1:])...)
	}
	return
}
