package shipping

import "fmt"

type Stack struct {
	MaxLength    int // -1 signifies no max length
	CurrentIndex int
	Items        []rune
}

func (s Stack) String() string {
	return fmt.Sprintf("[%v,%v,%v]", s.MaxLength, s.Items, string(s.Items))
}

// Create a new stack with the given max length and (optional) runes.
// A max length of -1 represents no limit.
func NewStack(maxLength int, chars ...rune) *Stack {
	if maxLength != -1 && maxLength < len(chars) {
		panic(
			fmt.Sprint("cannot create stack with maxLength", maxLength, "using", len(chars), "items", chars),
		)
	}

	s := &Stack{
		MaxLength:    maxLength,
		CurrentIndex: len(chars) - 1,
		Items:        chars,
	}

	return s
}

func (s *Stack) Push(r rune) {
	// if the struct has a max length & it's at its limit,
	// return a panic overflow
	if (s.MaxLength != -1) && (s.MaxLength == s.CurrentIndex+1) {
		panic(fmt.Sprintf("stack overflow adding %v to stack %+v", string(r), s))
	}

	fmt.Println("Pushing to index:", s.CurrentIndex+1, " rune ", r)
	s.CurrentIndex += 1
	s.Items = append(s.Items, r)
}

// Pop off the last item of the stack and return it.
func (s *Stack) Pop() rune {
	// panic if the stack is empty
	if s.CurrentIndex == -1 {
		panic("attempted to Pop() empty stack!")
	}

	r := s.Items[s.CurrentIndex]
	s.Items = s.Items[:s.CurrentIndex]
	s.CurrentIndex -= 1

	return r
}
