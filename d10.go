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

type coordinate struct {
	x     int
	y     int
	value int
}

func traverseMap(total [][]int, neighbours []coordinate, loc coordinate, foundNines map[coordinate]int) int {
	sum := 0
	for _, v := range neighbours {
		if loc.value == 8 && v.value == 9 {
			foundNines[v] = foundNines[v] + 1
			sum++
		}
		if loc.value+1 == v.value && loc.value != 8 {
			sum += traverseMap(total, getNeighbours(total, v), v, foundNines)
		}
	}
	return sum
}

func getNeighbours(total [][]int, loc coordinate) []coordinate {
	locs := make([]coordinate, 0)
	direction := []int{-1, 1}
	for _, v := range direction {
		if loc.x+v >= len(total[0]) || loc.x+v < 0 {
			continue
		}
		locs = append(locs, coordinate{x: loc.x + v, y: loc.y, value: total[loc.y][loc.x+v]})
	}

	for _, v := range direction {
		if loc.y+v >= len(total) || loc.y+v < 0 {
			continue
		}
		locs = append(locs, coordinate{x: loc.x, y: loc.y + v, value: total[loc.y+v][loc.x]})
	}

	return locs
}
