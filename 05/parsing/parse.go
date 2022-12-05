package parsing

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/cdl/advent_of_code_2022/05/shipping"
	"github.com/cdl/advent_of_code_2022/05/utils"
)

// Parse out an input string into a list of Stacks and Instructions.
func ParseInputFile(in string) []string {
	t := utils.ReadFileString(in)
	parts := strings.Split(t, "\n\n")

	header := strings.Split(parts[0], "\n")
	// instructions := strings.Split(parts[1], "\n")

	// regex, err := regexp.Compile(`\[([A-Z])\]`)
	// utils.Check(err)

	// first string always contains trailing whitespace
	// for empty columns, so we can divide its length by 4
	// to determine how many columns to parse out
	// ! this sucks and i hate it but i'm sick of spending time on this so idc anymore it works
	rows := len(header)
	columns := int(math.Ceil(float64(len(header[0])) / float64(4)))
	fmt.Printf("found %v rows\t", rows-1)
	fmt.Printf("found %v columns\n", columns)

	var stacks []*shipping.Stack

	// Instead of splitting on whitespace and properly parsing out [A] [B] [C] tokens
	// I'm opting to rely on consistency in whitespacing etc in the input and am using
	// character offsets to substring here in order to find the character.
	for x := 0; x < columns; x++ {
		var chars []rune
		for y := 0; y < rows-1; y++ {
			fmt.Printf("finding character for row %v col %v\n", x+1, y+1)
			line := header[y]
			char, _ := utf8.DecodeRuneInString(line[3*x+x+1 : 3*x+x+2])

			if char == ' ' {
				continue
			}

			chars = append(chars, char)
		}

		stacks = append(stacks, shipping.NewStack(-1, chars...))
	}

	fmt.Println(stacks)

	return []string{}
}
