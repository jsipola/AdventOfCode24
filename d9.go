package adventofcode24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day9() {
	data := ParseInputData("data/d9.txt")

	allsumsPart1 := 0
	allsums := 0
	blockSlices := make([]string, 0)
	input := strings.Split(data[0], "")
	for index, v := range input {
		num, _ := strconv.Atoi(v)
		if index%2 == 1 {
			blockSlices = append(blockSlices, slices.Repeat([]string{"."}, num)...)
		} else {
			blockSlices = append(blockSlices, slices.Repeat([]string{strconv.Itoa(index / 2)}, num)...)
		}
	}

	final := make([]string, 0)
	moveblocks(blockSlices, &final)
	for index, v := range final {
		val, _ := strconv.Atoi(v)
		allsumsPart1 += index * val
	}
	fmt.Println("Day 8 Part 1:    ", allsumsPart1)
	fmt.Println("Day 8 Part 2:    ", allsums)
}
