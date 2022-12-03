package main

import (
	"fmt"

	"github.com/DanielPettersson/aoc2022/common"
	"github.com/DanielPettersson/aoc2022/puzzle"
)

func getPrio(item rune) int {
	if item < 97 {
		return int(item) - 38
	} else {
		return int(item) - 96
	}
}

func main() {
	lines, err := puzzle.ReadInput("../input.txt")
	if err != nil {
		panic(err)
	}
	var sumOfPrios int

	for _, line := range lines {
		halfLineLen := len(line) / 2

		firstCompartment := line[0:halfLineLen]
		secondCompartment := line[halfLineLen:]

		firstCompartmentSet := common.SetOfChars(firstCompartment)
		secondCompartmentSet := common.SetOfChars(secondCompartment)

		commonSet := common.Intersection(firstCompartmentSet, secondCompartmentSet)

		if len(commonSet) != 1 {
			panic(commonSet)
		}

		var commonItem rune
		for k := range commonSet {
			commonItem = k
			break
		}

		sumOfPrios += getPrio(commonItem)
	}
	fmt.Println(sumOfPrios)
}
