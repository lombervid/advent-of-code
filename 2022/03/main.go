package main

import (
	"defaults"
	"fmt"
	"utils"
)

const day = "03"

func main() {
	fmt.Println("Day", day)

	filepath := utils.GetInputFilePath(defaults.Year, day)
	fileinfo := utils.ReadFile(filepath)
	rucksacks := getRucksacks(fileinfo)
	repeated := findRepeated(rucksacks)
	sumPriorities := getPriorities(repeated)

	fmt.Println("Part 1:", sumPriorities)
}

func getPriorities(s []string) int {
	var total int

	for _, value := range s {
		total += prioritize(value)
	}

	return total
}

func prioritize(s string) int {
	if s >= "a" && s <= "z" {
		return int([]rune(s)[0]) - 96
	}

	return int([]rune(s)[0]) - 38
}

func findRepeated(rucksacks [2][]map[string]int) []string {
	repeated := make([]string, 0)

	for i, row := range rucksacks[0] {
		for rsType := range row {
			if _, exists := rucksacks[1][i][rsType]; exists {
				repeated = append(repeated, rsType)
				break
			}
		}
	}

	return repeated
}

func getRucksacks(info []string) [2][]map[string]int {
	rucksacks := [2][]map[string]int{}

	for _, value := range info {
		length := len(value) / 2
		rs1 := value[:length]
		rs2 := value[length:]

		sack1 := make(map[string]int, length)
		sack2 := make(map[string]int, length)

		for i := 0; i < length; i++ {
			sack1[string(rs1[i])] = 1
			sack2[string(rs2[i])] = 1
		}

		rucksacks[0] = append(rucksacks[0], sack1)
		rucksacks[1] = append(rucksacks[1], sack2)
	}

	return rucksacks
}
