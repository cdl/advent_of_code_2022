package main

//
// Advent of Code 2022 - Day 2
//
// $ go run main.go
// # assumes the AOC input file is located alongside the source as input.txt
//

import (
	"fmt"
	"strings"

	"github.com/cdl/advent/02/utils"
)

// Choice represents what action a player chose to take for a given round.
type Choice int

const (
	Invalid  Choice = -1
	Rock     Choice = 0
	Paper    Choice = 1
	Scissors Choice = 2
)

// Parse a Choice input string as a properly typed enum value.
func ParseChoice(in string) Choice {
	switch in {
	case "A":
		return Rock
	case "X":
		return Rock
	case "B":
		return Paper
	case "Y":
		return Paper
	case "C":
		return Scissors
	case "Z":
		return Scissors
	}

	return Invalid
}

type Result int

const (
	Lose Result = 0
	Draw Result = 1
	Win  Result = 2
)

func ParseResult(in string) Result {
	switch in {
	case "X":
		return Lose
	case "Z":
		return Win
	default:
		return Draw
	}
}

// Round represents an individual "game", comprised of two choices and an outcome.
type Round struct {
	A      Choice
	B      Choice
	Result Result
	Score  int
}

// For a given string like `A X`, parse it out as an instance of a Round.
func ParseRound(in string) *Round {
	round := strings.Split(in, " ")

	r := &Round{
		A:      ParseChoice(round[0]),
		Result: ParseResult(round[1]),
	}

	r.B = CalculateMove(r.A, r.Result)
	r.Score = r.CalculateScore()

	return r
}

func CalculateMove(choice Choice, result Result) Choice {
	switch result {
	case Win:
		switch choice {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}

	case Draw:
		return choice

	case Lose:
		switch choice {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}

	// otherwise, return -1
	return -1
}

// Taken from:
// https://learningpenguin.net/2020/02/06/a-simple-algorithm-for-calculating-the-result-of-rock-paper-scissors-game/
func CalculateResult(a Choice, b Choice) Result {
	if (a+1)%3 == b {
		return Win
	} else if a == b {
		return Draw
	} else {
		return Lose
	}
}

// Calculate the score for the round.
func (r *Round) CalculateScore() int {
	score := 0

	// Determine score for pick.
	if r.B == Rock {
		score += 1
	} else if r.B == Paper {
		score += 2
	} else if r.B == Scissors {
		score += 3
	}

	// Determine win/draw bonus.
	if (r.A+1)%3 == r.B {
		// Player B won
		score += 6
	} else if r.A == r.B {
		// It is a draw
		score += 3
	}

	return score
}

func main() {
	input := utils.ReadFileString("input.txt")
	inputRounds := strings.Split(input, "\n")

	total := 0
	scores := []int{}

	for _, i := range inputRounds {
		r := ParseRound(i)
		scores = append(scores, r.Score)
		total += r.Score
	}

	fmt.Println("Total: ", total)
	fmt.Println("Rounds: ", len(scores))
}
