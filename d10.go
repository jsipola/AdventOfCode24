package adventofcode24

import (
	"fmt"
)

func day10() {
	data := ParseInputData("data/d10.txt")

	allsumsPart1 := 0
	allsumsPart2 := 0
	allsums := 0
	allInts := make([][]int, len(data))
	for index, v := range data {
		allInts[index] = convertToInts(v)
	}

	for yIndex, v := range allInts {
		for xIndex, x := range v {
			if x == 0 {
				nines := make(map[coordinate]int, 0)
				allsumsPart1 += traverseMap(allInts, getNeighbours(allInts, coordinate{x: xIndex, y: yIndex, value: x}), coordinate{x: xIndex, y: yIndex, value: x}, nines)
				allsums += len(nines)
				for _, v := range nines {
					allsumsPart2 += v
				}
			}
		}
	}

	fmt.Println("Day 10 Part 1:    ", allsums)
	fmt.Println("Day 10 Part 2:    ", allsumsPart2)
}
