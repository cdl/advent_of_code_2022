package shipping

import "fmt"

type Stack struct {
	MaxLength    int // -1 signifies no max length
	CurrentIndex int
	Items        []rune
}

func (s Stack) String() string {
	return fmt.Sprintf("[%v]", string(s.Items))
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
		CurrentIndex: -1,
	}

	// Ideally we'd reverse this inline here but I'll just manually push to create
	// new stacks to get the same effect.
	for x := range chars {
		s.Push(chars[len(chars)-1-x])
	}

	return s
}

func (s *Stack) Push(r rune) {
	// if the struct has a max length & it's at its limit,
	// return a panic overflow
	if (s.MaxLength != -1) && (s.MaxLength == s.CurrentIndex+1) {
		panic(fmt.Sprintf("stack overflow adding %v to stack %+v", string(r), s))
	}

	// fmt.Println("Pushing to index:", s.CurrentIndex+1, " rune ", r)
	s.CurrentIndex += 1
	s.Items = append(s.Items, r)
}

func (s *Stack) PushMany(chars ...rune) {
	// if the struct has a max length & it's at its limit,
	// return a panic overflow
	if (s.MaxLength != -1) && (s.MaxLength == s.CurrentIndex+len(chars)) {
		panic(fmt.Sprintf("stack overflow adding %v to stack %+v", string(chars), s))
	}

	s.CurrentIndex += len(chars)
	s.Items = append(s.Items, chars...)
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

// Pop an arbitrary amount of items off the top of the stack, preserving their order.
func (s *Stack) PopMany(quantity int) []rune {
	s.CurrentIndex -= quantity
	popped := s.Items[len(s.Items)-quantity : len(s.Items)]
	s.Items = s.Items[:len(s.Items)-quantity]

	return popped
}

func (s *Stack) Peek() rune {
	// panic if the stack is empty
	if s.CurrentIndex == -1 {
		panic("attempted to Peek() an empty stack!")
	}

	return s.Items[s.CurrentIndex]
}
