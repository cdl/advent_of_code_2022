package main

import (
	"fmt"

	"github.com/cdl/advent_of_code_2022/04/assignments"
	"github.com/cdl/advent_of_code_2022/04/utils"
)

func main() {
	in := utils.ReadFileString("input")
	parsed := assignments.ParseAssignmentList(in)
	var eitherWithin []bool

	for _, i := range parsed {
		if i[0].IsWithin(i[1]) || i[1].IsWithin(i[0]) {
			eitherWithin = append(eitherWithin, true)
		}
	}

	fmt.Println(len(eitherWithin))

}
