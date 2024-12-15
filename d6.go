package adventofcode24

import (
	"fmt"
	"strings"
)

func day6() {
	data := ParseInputData("data/d6.txt")

	sumsPart2 := 0

	directionMap := map[string]Path{
		"<": {yDir: 0, xDir: -1, Next: "^"},
		">": {yDir: 0, xDir: 1, Next: "v"},
		"^": {yDir: -1, xDir: 0, Next: ">"},
		"v": {yDir: 1, xDir: 0, Next: "<"},
	}
	LocMap := make(map[Loc]int, 0)

	/* y, x := findCurrentPos(data)
	dire := string(data[y][x])

	fmt.Println(y, x)
	fmt.Println(directionMap[dire]) */

	locs := make([]Loc, 0)
	sqLocs := make([]Loc, 0)
	traverse := true
	for traverse {
		startY, startX := findCurrentPos(data)
		dire := string(data[startY][startX])

		rowToUpdate := data[startY]
		//cpyRow := strings.Clone(rowToUpdate)
		//fmt.Println(rowToUpdate)
		updatedRow := strings.Replace(rowToUpdate, dire, ".", 1)
		y, x, locs2 := findPath(directionMap, data)
		for _, v := range locs2 {
			LocMap[v] = 1
		}
		// Part 2
		sqLocs = append(sqLocs, Loc{x: x, y: y})

		/* 		if len(sqLocs) > 3 {
			// check paths? and compare
			for _, v := range locs2 {
				tmpData := slices.Clone(data)
				tmpUpdatedRow := strings.Replace(tmpData[startY], dire, ".", 1)
				tmpData[startY] = tmpUpdatedRow
				tmpData[v.y] = JoinUpdatedRow(tmpData[v.y], directionMap[dire].Next, v.x, v.y)
				fmt.Println(tmpData[v.y])
				a, _ := traversePath(tmpData, v, []Loc{})
				if !a {
					sumsPart2++
					//fmt.Println(tmpData[v.y])
					//fmt.Println("traverse result:", a, "For x:", v.x, "y:", v.y)
				}

				//fmt.Println()
			}
			//break
			//fmt.Println(sqLocs)
			//fmt.Println(locs2)
		} */
		/* if len(sqLocs) == 4 {
			sqLocs = sqLocs[1:]
		} */
		locs = append(locs, locs2...)
		if y == -1 && x == -1 {
			traverse = false
			break
		}

		// remove old guard Char
		data[startY] = updatedRow
		// update new guard Char
		data[y] = JoinUpdatedRow(data[y], directionMap[dire].Next, x, y)

		//fmt.Println("x:", x, "y:", y)
	}
	/* 	for _, v := range data {
		fmt.Println(v)
	} */
	/* 	fmt.Println(len(LocMap))
	   	fmt.Println(len(sqLocs)) */
	/* 	slices.SortFunc(locs, func(a Loc, b Loc) int {
	   		if a.x == b.x && a.y == b.y {
	   			return 0
	   		}
	   		if a.x < b.x && a.y == b.y {
	   			return 1
	   		}
	   		return -1
	   	})
	   	//fmt.Println(locs)
	   	newlocs := slices.CompactFunc(locs, func(a Loc, b Loc) bool {
	   		return a.x == b.x && a.y == b.y
	   	}) */
	/* 	for _, v := range newlocs {
		fmt.Println(v)
	} */

	fmt.Println("Day 6 Part 1: ", len(LocMap))
	fmt.Println("Day 6 Part 2: ", sumsPart2)
}

func day6Part2() {
	data := ParseInputData("data/d6.txt")

	sumsPart2 := 0

	directionMap := map[string]Path{
		"<": {yDir: 0, xDir: -1, Next: "^"},
		">": {yDir: 0, xDir: 1, Next: "v"},
		"^": {yDir: -1, xDir: 0, Next: ">"},
		"v": {yDir: 1, xDir: 0, Next: "<"},
	}
	LocMap := make(map[Loc]int, 0)
	LocMap2 := make(map[Loc]int, 0)

	/* y, x := findCurrentPos(data)
	dire := string(data[y][x])

	fmt.Println(y, x)
	fmt.Println(directionMap[dire]) */

	locs := make([]Loc, 0)
	sqLocs := make([]Loc, 0)
	//traverse := true
	startY, startX := findCurrentPos(data)
	dire := string(data[startY][startX])
	startpath := directionMap[dire]
	_, paths := traversePath(data, Loc{x: startX, y: startY}, startpath, []Loc{})
	//fmt.Println(startpath)
	//fmt.Println(paths)

	for _, v := range paths {
		// Remove direction incase been to same place more than once
		v = Loc{x: v.x, y: v.y}
		LocMap2[v] = 1
	}

	for index, v := range paths {

		y, x, locs2 := findBlocker(v.x, v.y, startpath, data)
		startpath = directionMap[v.direction]
		//fmt.Println(startpath)
		for _, v := range locs2 {
			LocMap[v] = 1
		}
		// Part 2
		//fmt.Println("Start travelsar for sub blockers", v.x, v.y)
		//fmt.Println(locs2[0].x, locs2[0].y)
		a, _ := traversePath(data, v, startpath, paths[:index])
		//fmt.Println(a)
		if !a {
			//fmt.Println(new)
			sumsPart2++
		}
		sqLocs = append(sqLocs, Loc{x: x, y: y})
		//break
		//fmt.Println(locs2)

		/* 		if len(sqLocs) > 3 {
			// check paths? and compare
			for _, v := range locs2 {
				tmpData := slices.Clone(data)
				tmpUpdatedRow := strings.Replace(tmpData[startY], dire, ".", 1)
				tmpData[startY] = tmpUpdatedRow
				tmpData[v.y] = JoinUpdatedRow(tmpData[v.y], directionMap[dire].Next, v.x, v.y)
				fmt.Println(tmpData[v.y])
				a := traversePath(tmpData, v, []Loc{})
				if !a {
					sumsPart2++
					//fmt.Println(tmpData[v.y])
					//fmt.Println("traverse result:", a, "For x:", v.x, "y:", v.y)
				}

				//fmt.Println()
			}
			//break
			//fmt.Println(sqLocs)
			//fmt.Println(locs2)
		} */
		/* if len(sqLocs) == 4 {
			sqLocs = sqLocs[1:]
		} */
		locs = append(locs, locs2...)
		/* 		if y == -1 && x == -1 {
			//traverse = false
			break
		} */

		// remove old guard Char
		//data[startY] = updatedRow
		// update new guard Char
		//data[y] = JoinUpdatedRow(data[y], directionMap[dire].Next, x, y)

		//fmt.Println("x:", x, "y:", y)
	}
	/* 	for _, v := range data {
		fmt.Println(v)
	} */
	/* 	fmt.Println(len(LocMap))
	   	fmt.Println(len(sqLocs)) */
	/* 	slices.SortFunc(locs, func(a Loc, b Loc) int {
	   		if a.x == b.x && a.y == b.y {
	   			return 0
	   		}
	   		if a.x < b.x && a.y == b.y {
	   			return 1
	   		}
	   		return -1
	   	})
	   	//fmt.Println(locs)
	   	newlocs := slices.CompactFunc(locs, func(a Loc, b Loc) bool {
	   		return a.x == b.x && a.y == b.y
	   	}) */
	/* 	for _, v := range newlocs {
		fmt.Println(v)
	} */

	fmt.Println("Day 6 Part 1: ", len(LocMap2))
	fmt.Println("Day 6 Part 2: ", sumsPart2)
}
