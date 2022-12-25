package main

import (
	"2022/defaults"
	"fmt"
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

	fmt.Println(calories)
}

func getElvesCalories(info []string) map[int][]int64 {
	var elf int = 1
	elves := make(map[int][]int64, 0)

	for _, value := range info {
		if value == "" {
			elf++
			continue
		}
		calories, err := strconv.ParseInt(value, 10, 64)
		utils.CheckErrNil(err)
		elves[elf] = append(elves[elf], calories)
	}

	return elves
}

func getMaxCalories(m map[int][]int64) int64 {
	var maxCalories int64

	for _, calories := range m {
		var totalCalories int64
		for _, calorie := range calories {
			totalCalories += calorie
		}

		if totalCalories > maxCalories {
			maxCalories = totalCalories
		}
	}

	return maxCalories
}
