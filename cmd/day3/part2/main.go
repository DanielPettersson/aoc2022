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

	for i := range lines {
		if i%3 != 0 {
			continue
		}

		firstElf := common.SetOfChars(lines[i])
		secondElf := common.SetOfChars(lines[i+1])
		thirdElf := common.SetOfChars(lines[i+2])

		commonSet := common.Intersection(common.Intersection(firstElf, secondElf), thirdElf)

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
