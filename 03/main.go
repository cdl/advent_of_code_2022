package main

//
// Advent of Code 2022 - Day 3
//
// $ go run main.go
// # assumes the AOC input file is located alongside the source as input.txt
//

import (
	"fmt"
	"strings"

	"github.com/cdl/advent2022/03/rucksack"
	"github.com/cdl/advent2022/03/utils"
)

func main() {
	// in := utils.ReadFileString("control.txt")
	in := utils.ReadFileString("input.txt")
	// in := utils.ReadFileString("input2.txt")
	sacks := strings.Split(in, "\n")
	itemPriorities := rucksack.GeneratePriorityMap()

	var parsed []rucksack.Rucksack

	for _, str := range sacks {
		r := rucksack.ParseRucksack(itemPriorities, str)
		parsed = append(parsed, *r)
	}

	// priorities := []int{}
	sum := 0

	for i := 0; i < len(parsed); i += 3 {
		// Compare three rucksacks at a time to determine the common item.
		one := parsed[i]
		two := parsed[i+1]
		three := parsed[i+2]

		common := rucksack.FindCommonItem(one, two, three)

		// priorities = append(priorities, itemPriorities[common])
		sum += itemPriorities[common]
	}

	// fmt.Println(priorities)
	fmt.Println(sum)
}
