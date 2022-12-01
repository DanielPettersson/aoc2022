package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/DanielPettersson/aoc2022/file"
)

func main() {
	lines, err := file.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var elfCaloryList sort.IntSlice
	var elfCalorySum int

	for _, line := range lines {

		if line == "" {
			elfCaloryList = append(elfCaloryList, elfCalorySum)
			elfCalorySum = 0
		} else {
			num, _ := strconv.Atoi(line)
			elfCalorySum += num
		}
	}
	sort.Sort(sort.Reverse(elfCaloryList))

	fmt.Println(elfCaloryList[0] + elfCaloryList[1] + elfCaloryList[2])
}
