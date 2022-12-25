package main

import (
	"defaults"
	"fmt"
	"strings"
	"utils"
)

const day = "02"

func main() {
	fmt.Println("Day", day)
	inputFile := utils.GetInputFilePath(defaults.Year, day)

	fileinfo := utils.ReadFile(inputFile)

	score := getTotalScorePart1(fileinfo)
	fmt.Println("Part 1:", score)
}

func getTotalScorePart1(info []string) int {
	var totalScore int

	shape := map[string]map[string]int{
		"X": {"A": 3, "B": 0, "C": 6},
		"Y": {"A": 6, "B": 3, "C": 0},
		"Z": {"A": 0, "B": 6, "C": 3},
	}

	for _, value := range info {
		scores := strings.Split(value, " ")
		switch scores[1] {
		case "X":
			totalScore += 1 + shape[scores[1]][scores[0]]
		case "Y":
			totalScore += 2 + shape[scores[1]][scores[0]]
		case "Z":
			totalScore += 3 + shape[scores[1]][scores[0]]
		}
	}

	return totalScore
}
