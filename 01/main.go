package main

//
// Advent of Code 2022 - Day 1
//
// $ go run main.go
// # assumes the AOC input file is located alongside the source as `input`
//

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Given a raw input file, parse out the individual elf inventories.
func parseInventories(in string) []string {
	return strings.Split(string(in), "\n\n")
}

// Parse an elves inventory string into a list of calorie values.
func parseInventory(inv string) []int {
	items := strings.Split(inv, "\n")

	var parsedItems []int
	for _, item := range items {
		val, err := strconv.Atoi(item)
		if err == nil {
			parsedItems = append(parsedItems, val)
		}
	}

	return parsedItems
}

// Given an elves' raw inventory, calculate its total.
func getInventoryTotal(inv string) int {
	total := 0
	items := parseInventory(inv)

	for _, i := range items {
		total += i
	}

	return total
}

func main() {
	input, err := os.ReadFile("input")
	// input, err := os.ReadFile("control-input")
	check(err)

	inventories := parseInventories(string(input))

	totals := []int{}
	for _, inv := range inventories {
		total := getInventoryTotal(inv)
		totals = append(totals, total)
	}

	// Sort the inventory totals to determine the highest, then pluck the top three.
	sort.Ints(totals)
	topThree := totals[len(totals)-3:]
	topThreeTotal := 0

	// Finally, sum those top three inventories, and output the result.
	for _, i := range topThree {
		topThreeTotal += i
	}

	fmt.Println(topThreeTotal)
}
