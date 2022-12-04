package assignments

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cdl/advent_of_code_2022/04/utils"
)

type Assignment struct {
	Start int
	End   int
}

func (a Assignment) String() string {
	return fmt.Sprintf("[%v-%v]", a.Start, a.End)
}

// Check to see if the current assignment is contained within the given one.
func (a *Assignment) IsWithin(b *Assignment) bool {
	if a.Start >= b.Start && a.End <= b.End {
		return true
	}

	return false
}

// For a given "assignment" string (looks like `1-3`), parse out the start
// and end section numbers and create a new Assignment for it.
func ParseAssignment(in string) *Assignment {
	sections := strings.Split(in, "-")

	start, err := strconv.Atoi(sections[0])
	utils.Check(err)

	end, err := strconv.Atoi(sections[1])
	utils.Check(err)

	return &Assignment{
		Start: start,
		End:   end,
	}
}

// For a given "assignment pair" string (looks like `1-3,4-6`), parse out the
// individual assignments. Returns an array containing the pair of assignments.
func ParseAssignmentPair(in string) []*Assignment {
	assignments := strings.Split(in, ",")
	var parsed []*Assignment
	for _, a := range assignments {
		parsed = append(parsed, ParseAssignment(a))
	}

	return parsed
}

// Returns an array of arrays of assignments.
func ParseAssignmentList(in string) [][]*Assignment {
	lines := strings.Split(in, "\n")
	var pairs [][]*Assignment

	for _, line := range lines {
		pairs = append(pairs, ParseAssignmentPair(line))
	}

	return pairs
}
