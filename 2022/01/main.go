package main

import (
	"defaults"
	"fmt"
	"sort"
	"strconv"
	"utils"
)

const day = "01"

func main() {
	fmt.Println("Day", day)
	inputFile := utils.GetInputFilePath(defaults.Year, day)

	fileinfo := utils.ReadFile(inputFile)

	elves := getElvesCalories(fileinfo)
	calories := getMaxCalories(elves)

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	fmt.Println("Part 1:", calories[0])
	fmt.Println("Part 2:", calories[0]+calories[1]+calories[2])
}

func getElvesCalories(info []string) map[int][]int {
	var elf int = 1
	elves := make(map[int][]int, 0)

	for _, value := range info {
		if value == "" {
			elf++
			continue
		}
		calories, err := strconv.ParseInt(value, 10, 64)
		utils.CheckErrNil(err)
		elves[elf] = append(elves[elf], int(calories))
	}

	return elves
}

func getMaxCalories(m map[int][]int) []int {
	var maxCalories []int

	for _, calories := range m {
		var totalCalories int
		for _, calorie := range calories {
			totalCalories += calorie
		}

		maxCalories = append(maxCalories, totalCalories)
	}

	return maxCalories
}
