package adventofcode24

import (
	"fmt"
	"strings"
)

func day4() {
	data := ParseInputData("data/d4.txt")

	directions1 := map[string][]int{"NW": {-1, -1}, "N": {-1, 0}, "NE": {-1, 1}, "E": {0, 1}, "SE": {1, 1}, "S": {1, 0}, "SW": {1, -1}, "W": {0, -1}}
	directions2 := map[string][]int{"SE": {1, 1}, "SW": {1, -1}}
	sums := 0
	sums2 := 0
	type Loc struct {
		x int
		y int
	}
	foundXs := make(map[Loc][]string, 0)
	for yIndex, v := range data {
		for xIndex, char := range v {
			if char == 'X' {
				for _, direction := range directions1 {
					yDir := direction[0]
					xDir := direction[1]
					str := make([]string, 0)
					for i := 0; i < 4; i++ {
						yloc := yIndex + (yDir * i)
						xloc := xIndex + (xDir * i)
						if yloc < 0 || yloc >= len(data) || xloc < 0 || xloc >= len(v) {
							break
						}
						str = append(str, string(data[yloc][xloc]))
					}
					if len(str) != 4 {
						continue
					}
					if strings.Join(str, "") == "XMAS" {
						sums++
					}
				}
			}
			// Part2
			if char == 'M' || char == 'S' {
				for orientation, direction := range directions2 {
					yDir := direction[0]
					xDir := direction[1]
					str := make([]string, 0)
					for i := 0; i < 3; i++ {
						yloc := yIndex + (yDir * i)
						xloc := xIndex + (xDir * i)
						if yloc < 0 || yloc >= len(data) || xloc < 0 || xloc >= len(v) {
							break
						}
						str = append(str, string(data[yloc][xloc]))
					}
					if len(str) != 3 {
						continue
					}
					joinedStr := strings.Join(str, "")
					if joinedStr == "MAS" || joinedStr == "SAM" {
						// Find the A location
						Aloc := Loc{x: xIndex + xDir, y: yIndex + yDir}
						foundXs[Aloc] = append(foundXs[Aloc], orientation)
					}
				}
			}
		}
	}

	fmt.Println("Day 4 Part 1: ", sums)
	for _, v := range foundXs {
		// Check A location has two directions
		if len(v) == 2 {
			sums2++
		}
	}
	fmt.Println("Day 4 Part 2: ", sums2)
}
