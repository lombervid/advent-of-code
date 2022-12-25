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
	options := getPlayersSelections(fileinfo)

	score := getTotalScorePart1(options)
	fmt.Println("Part 1:", score)

	score = getTotalScore(options)
	fmt.Println("Part 2:", score)
}

func getPlayersSelections(info []string) [][]string {
	options := make([][]string, 0)

	for _, value := range info {
		options = append(options, strings.Split(value, " "))
	}

	return options
}

func getTotalScore(matches [][]string) int {
	var totalScore int

	shapeToUse := map[string]map[string]string{
		"A": {"X": "C", "Y": "A", "Z": "B"},
		"B": {"X": "A", "Y": "B", "Z": "C"},
		"C": {"X": "B", "Y": "C", "Z": "A"},
	}

	rewards := map[string]map[string]int{
		"A": {"A": 3, "B": 0, "C": 6},
		"B": {"A": 6, "B": 3, "C": 0},
		"C": {"A": 0, "B": 6, "C": 3},
	}

	for _, match := range matches {
		mySelection := shapeToUse[match[0]][match[1]]
		totalScore += int([]rune(mySelection)[0] - 64) // A => 65; B => 66; C => 67
		totalScore += rewards[mySelection][match[0]]
	}

	return totalScore
}

func getTotalScorePart1(matches [][]string) int {
	var totalScore int

	rewards := map[string]map[string]int{
		"X": {"A": 3, "B": 0, "C": 6},
		"Y": {"A": 6, "B": 3, "C": 0},
		"Z": {"A": 0, "B": 6, "C": 3},
	}

	for _, option := range matches {
		totalScore += int([]rune(option[1])[0] - 87) // X => 88; Y => 89; Z => 90
		totalScore += rewards[option[1]][option[0]]
	}

	return totalScore
}
