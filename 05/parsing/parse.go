package parsing

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
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

	// Instead of splitting on whitespace or using regex to parse out [A] [B] [C] tokens,
	// I'm opting to rely on consistency in whitespacing etc in the input and am using
	// character offsets to generate substrings.
	for x := 0; x < columns; x++ {
		var chars []rune
		for y := 0; y < rows-1; y++ {
			fmt.Printf("finding character for row %v col %v\n", x+1, y+1)
			line := header[y]
			char, _ := utf8.DecodeRuneInString(line[3*x+x+1 : 3*x+x+2])

			if char == ' ' {
				// skip any empty slots
				continue
			}

			chars = append(chars, char)
		}

		stacks = append(stacks, shipping.NewStack(-1, chars...))
	}

	// ! Instead of creating some form of object to contain a list of Stacks and
	// ! a Move() function, I'm opting to do it all inline here.

	// Finished assembling the stacks. Let's parse out the instructions.
	regex, err := regexp.Compile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
	utils.Check(err)

	instructions := regex.FindAllStringSubmatch(parts[1], -1)
	for _, i := range instructions {
		quantity, err := strconv.Atoi(i[1])
		utils.Check(err)

		from, err := strconv.Atoi(i[2])
		utils.Check(err)

		to, err := strconv.Atoi(i[3])
		utils.Check(err)

		fromIndex := from - 1
		toIndex := to - 1

		moved := stacks[fromIndex].PopMany(quantity)
		fmt.Printf("moving `%v` from %v to %v\n", string(moved), from, to)
		stacks[toIndex].PushMany(moved...)

		// for x := 0; x < quantity; x++ {
		// 	moved := stacks[fromIndex].Pop()

		// 	fmt.Printf("moving rune %v from %v to %v\n", string(moved), from, to)
		// 	stacks[toIndex].Push(moved)
		// }
	}

	fmt.Printf("Final state: %v\n", stacks)

	// lastly, output the rune on top of each stack
	for _, stack := range stacks {
		fmt.Print(string(stack.Peek()))
	}

	fmt.Println()

	return []string{}
}
