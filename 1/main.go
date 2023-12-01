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
	reg, _ := regexp.Compile("[0-9]")
	result := 0
	for _, line := range lines {
		matches := reg.FindAll(line, -1)
		cal, _ := strconv.Atoi(fmt.Sprintf("%s%s", matches[0], matches[len(matches)-1]))
		result += cal
	}
	fmt.Println(result)
}
