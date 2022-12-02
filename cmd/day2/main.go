package main

import (
	"fmt"

	"github.com/DanielPettersson/aoc2022/puzzle"
)

type shape int

const (
	rock    = 1
	paper   = 2
	scissor = 3
)

var (
	shapes = []shape{rock, paper, scissor}
)

func createShape(letter rune) (shape, error) {
	switch string(letter) {
	case "A":
		return rock, nil
	case "B":
		return paper, nil
	case "C":
		return scissor, nil
	}
	return -1, fmt.Errorf("invalid letter for shape: %v", letter)
}

func createOutcome(letter rune) (int, error) {
	switch string(letter) {
	case "X":
		return 0, nil
	case "Y":
		return 3, nil
	case "Z":
		return 6, nil
	}
	return -1, fmt.Errorf("invalid letter for outcome: %v", letter)
}

func (s shape) getOutcome(other shape) int {
	if s == other {
		return 3
	}
	if s == rock && other == scissor {
		return 6
	}
	if s == paper && other == rock {
		return 6
	}
	if s == scissor && other == paper {
		return 6
	}
	return 0
}

func main() {
	lines, err := puzzle.ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	var score int

	for _, line := range lines {
		runes := []rune(line)
		opponentShape, err := createShape(runes[0])
		if err != nil {
			panic(err)
		}
		wantedOutcome, err := createOutcome(runes[2])
		if err != nil {
			panic(err)
		}

		var mySelectedShape shape
		for _, myShape := range shapes {
			outcome := myShape.getOutcome(opponentShape)
			if outcome == wantedOutcome {
				mySelectedShape = myShape
				break
			}
		}

		score += wantedOutcome + int(mySelectedShape)

	}
	fmt.Println(score)
}
