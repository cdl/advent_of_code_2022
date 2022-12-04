package main

import (
	"fmt"

	"github.com/cdl/advent_of_code_2022/04/assignments"
	"github.com/cdl/advent_of_code_2022/04/utils"
)

func main() {
	in := utils.ReadFileString("input")
	parsed := assignments.ParseAssignmentList(in)
	var doesOverlap []bool

	for _, i := range parsed {
		if i[0].DoesOverlap(i[1]) || i[1].DoesOverlap(i[0]) {
			doesOverlap = append(doesOverlap, true)
		}
	}

	fmt.Println(len(doesOverlap))
}
