package adventofcode24

import (
	"fmt"
	"strings"
)

func day15() {
	data := ParseInputData("data/d15.txt")
	var startMap [][]string
	paths := make([]string, 0)
	var startPos Loc
	for index, v := range data {
		start := strings.Index(v, "@")
		if start != -1 {
			startPos = Loc{x: start, y: index}
			continue
		}
		if v == "" {
			tmp := data[:index]
			startMap = make([][]string, len(tmp))
			for index, v := range tmp {
				startMap[index] = append(startMap[index], strings.Split(v, "")...)
			}
			paths = data[index+1:]
			break
		}
	}

	for _, v := range paths {
		makeMoves(startMap, v, &startPos)
	}
	part1Sums := 0
	for yIndex, v := range startMap {
		for xIndex, v := range v {
			if v == "O" {
				part1Sums += 100*yIndex + xIndex
			}
		}
	}
	for _, v := range startMap {
		fmt.Println(v)
	}
	fmt.Println("Day 15 Part 1:", part1Sums)

}

func makeMoves(startMap [][]string, moves string, currentPos *Loc) {
	directionMap := map[string]Path{
		"<": {yDir: 0, xDir: -1},
		">": {yDir: 0, xDir: 1},
		"^": {yDir: -1, xDir: 0},
		"v": {yDir: 1, xDir: 0},
	}
	for _, v := range strings.Split(moves, "") {
		newPos := Loc{x: currentPos.x + directionMap[v].xDir, y: currentPos.y + directionMap[v].yDir}
		a := startMap[newPos.y][newPos.x]
		// check for box
		if a == "O" {
			// check for wall behind sbox
			for i := 1; i < len(startMap); i++ {
				if newPos.x+directionMap[v].xDir*i > len(startMap[0])-1 || newPos.x+directionMap[v].xDir*i < 0 {
					break
				}
				if newPos.y+directionMap[v].yDir*i > len(startMap)-1 || newPos.y+directionMap[v].yDir*i < 0 {
					break
				}
				if startMap[newPos.y+directionMap[v].yDir*i][newPos.x+directionMap[v].xDir*i] == "." {
					// save coordinates to swap box and move
					startMap[newPos.y+directionMap[v].yDir*i][newPos.x+directionMap[v].xDir*i] = "O"
					startMap[currentPos.y][currentPos.x] = "."
					startMap[newPos.y][newPos.x] = "@"
					*currentPos = newPos
					break
				}
				if startMap[newPos.y+directionMap[v].yDir*i][newPos.x+directionMap[v].xDir*i] == "#" {
					break
				}
			}
			continue
		}
		if a == "#" {
			// Do nothing
			continue
		}
		startMap[currentPos.y][currentPos.x] = "."
		startMap[newPos.y][newPos.x] = "@"
		*currentPos = newPos
	}
}
