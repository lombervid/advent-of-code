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
	dividedRucksacks := getDividedRucksacks(fileinfo)

	repeated := findRepeated(dividedRucksacks)
	sumPriorities := getPriorities(repeated)
	fmt.Println("Part 1:", sumPriorities)

	rucksacks := getRucksacks(fileinfo)
	badges := findBadges(rucksacks)
	badgesPriorities := getPriorities(badges)
	fmt.Println("Part 2:", badgesPriorities)
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

func findBadges(rucksacks []map[string]int) []string {
	badges := make([]string, 0)

	for i := 0; i < len(rucksacks); i += 3 {
		for item := range rucksacks[i] {
			if _, exists := rucksacks[i+1][item]; !exists {
				continue
			}

			if _, exists := rucksacks[i+2][item]; !exists {
				continue
			}

			badges = append(badges, item)
			break
		}
	}

	return badges
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

func getRucksacks(info []string) []map[string]int {
	rucksacks := []map[string]int{}

	for _, line := range info {
		length := len(line)
		sack := make(map[string]int, length)

		for _, item := range line {
			sack[string(item)] = 1
		}

		rucksacks = append(rucksacks, sack)
	}

	return rucksacks
}

func getDividedRucksacks(info []string) [2][]map[string]int {
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
