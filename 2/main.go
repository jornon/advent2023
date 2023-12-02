package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("2/input")
	lines := bytes.Split(data, []byte("\n"))
	a(lines)
	b(lines)
}

func b(lines [][]byte) {
	result := 0
	regex, _ := regexp.Compile("Game [0-9]+: ")
	for _, line := range lines {
		cubeMap := map[string]int{}
		game := regex.ReplaceAll(line, []byte(""))
		sets := bytes.Split(game, []byte(";"))
		for _, set := range sets {
			cubes := bytes.Split(set, []byte(","))
			for _, cubeColor := range cubes {
				parts := bytes.Split(bytes.TrimSpace(cubeColor), []byte(" "))
				color := string(parts[1])
				num, _ := strconv.Atoi(string(parts[0]))
				if v, ok := cubeMap[color]; ok {
					if num > v {
						cubeMap[color] = num
					}
				} else {
					cubeMap[color] = num
				}
			}
		}
		gameRes := 1
		for _, v := range cubeMap {
			gameRes *= v
		}
		result += gameRes
	}
	fmt.Printf("Result B: %d\n", result)
}

func a(lines [][]byte) {
	result := 0
	regex, _ := regexp.Compile("Game [0-9]+: ")
	for i, line := range lines {
		works := true
		game := regex.ReplaceAll(line, []byte(""))
		sets := bytes.Split(game, []byte(";"))
		for _, set := range sets {
			cubes := bytes.Split(set, []byte(","))
			for _, cubeColor := range cubes {
				if itWorks(cubeColor) == false {
					works = false
					break
				}
			}
			if !works {
				break
			}
		}
		if works {
			result += i + 1
		}
	}
	fmt.Printf("Result A: %d\n", result)
}

//only 12 red cubes, 13 green cubes, and 14 blue cubes

func itWorks(color []byte) bool {
	limitMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	parts := bytes.Split(bytes.TrimSpace(color), []byte(" "))
	if v, ok := limitMap[string(parts[1])]; ok {
		num, _ := strconv.Atoi(string(parts[0]))
		if num > v {
			//fmt.Printf("Too many %s cubes: %d\n", string(parts[1]), num)
			return false
		}
	} else {
		//fmt.Printf("Invalid color: %s\n", string(parts[1]))
		return false
	}
	return true
}
