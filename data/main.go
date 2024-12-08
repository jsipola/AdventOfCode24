package adventofcode24

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day1() {
	data := ParseInputData("./data/d1.txt")

	first := make([]float64, 0)
	second := make([]float64, 0)

	for _, v := range data {
		fields := strings.Fields(v)
		firstInt, _ := strconv.ParseFloat(fields[0], 64)
		secondInt, _ := strconv.ParseFloat(fields[1], 64)
		first = append(first, firstInt)
		second = append(second, secondInt)
	}
	slices.Sort(first)
	slices.Sort(second)
	sum := 0.0
	sum2 := 0.0
	for index, v := range first {
		// Part 1
		sum += math.Abs(v - second[index])

		// Part 2
		count := CountValues(second, v)
		sum2 += v * count
	}
	fmt.Printf("%f\n", sum)
	fmt.Printf("%f\n", sum2)
}

func day2() {
	data := ParseInputData("./data/d2.txt")

	safeCount := 0
rowLoop:
	for _, v := range data {
		fields := strings.Fields(v)
		intFields := make([]int, len(fields))
		for i, v := range fields {
			intFields[i], _ = strconv.Atoi(v)
		}
		isAscRow := 0
		for index, current := range intFields {
			if index+1 >= len(intFields) {
				continue
			}

			next := intFields[index+1]
			isValid, isAscPair := IsValidAndOrderedAsc(current, next)
			if isAscRow == 0 {
				isAscRow = isAscPair
			}
			if !isValid || isAscRow != isAscPair {
				continue rowLoop
			}
		}

		if slices.IsSorted(intFields) {
			safeCount++
		} else {
			slices.Reverse(intFields)
			if slices.IsSorted(intFields) {
				safeCount++
			}
		}
	}
	fmt.Printf("Day 2 Part 1: %d\n", safeCount)
}

func day2Part2() {
	data := ParseInputData("./data/d2.txt")

	safeCount := 0
	for _, v := range data {
		fields := strings.Fields(v)
		intFields := make([]int, len(fields))
		for i, v := range fields {
			intFields[i], _ = strconv.Atoi(v)
		}
		isValid := isDataValid(intFields)
		if isValid {
			safeCount++
			continue
		}
		for toRemoveIndex := range intFields {
			firstCpy := slices.Clone(intFields)
			firstRemoved := append(firstCpy[:toRemoveIndex], firstCpy[toRemoveIndex+1:]...)
			isValid = isDataValid(firstRemoved)
			if isValid {
				safeCount++
				break
			}
		}
	}
	fmt.Printf("Day 2 Part 2: %d\n", safeCount)
}

func day3(isPart2 bool) {
	data := ParseInputData("data/d3.txt")

	f := func(c rune) bool {
		return c == ')'
	}
	sums := 0
	enabled := true
	for _, v := range data {
		strs := strings.Split(v, "mul(")
		for index, v := range strs {
			v1 := strings.FieldsFunc(v, f)
			if len(v1) == 0 {
				continue
			}
			if index > 0 && isPart2 {
				dontIndex := strings.Index(strs[index-1], "don't()")
				doIndex := strings.Index(strs[index-1], "do()")
				if doIndex > dontIndex {
					enabled = true
				}
				if dontIndex > doIndex {
					enabled = false
				}
			}

			nums := strings.Split(v1[0], ",")
			if len(nums) != 2 {
				continue
			}
			mult1, err := strconv.Atoi(nums[0])
			if err != nil {
				continue
			}
			mult2, err := strconv.Atoi(nums[1])
			if err != nil {
				continue
			}
			if enabled {
				sums += mult1 * mult2
			}
		}
	}
	fmt.Printf("Day3 part ")
	if isPart2 {
		fmt.Print("2")
	} else {
		fmt.Print("1")
	}
	fmt.Printf(" %d\n", sums)
}

type Loc struct {
	x         int
	y         int
	direction string
}

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

func day5() {
	data := ParseInputData("data/d5.txt")

	sums := 0
	sumsPart2 := 0
	rules := make(map[int][]int, 0)

	for _, v := range data {
		v = strings.Trim(v, "\r\n")
		pages := strings.Split(v, "|")
		if len(v) == 0 {
			continue
		}
		if len(pages) == 2 {
			first, _ := strconv.Atoi(pages[0])
			second, _ := strconv.Atoi(pages[1])
			rules[first] = append(rules[first], second)
			continue
		}

		toUpdate := strings.Split(v, ",")
		updatePages := make([]int, len(toUpdate))
		for i, v := range toUpdate {
			updatePages[i], _ = strconv.Atoi(v)
		}

		if !slices.IsSortedFunc(updatePages, sortF(rules)) {
			slices.SortFunc(updatePages, sortF(rules))
			sumsPart2 = sumsPart2 + updatePages[((len(updatePages))/2)]
		} else {
			sums = sums + updatePages[((len(updatePages)-1)/2)]
		}
	}

	fmt.Println("Day 5 Part 1: ", sums)
	fmt.Println("Day 5 Part 2: ", sumsPart2)
}

type Path struct {
	yDir int
	xDir int
	Next string
}

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
	fmt.Println(startpath)
	fmt.Println(paths)

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
		fmt.Println("Start travelsar for sub blockers", v.x, v.y)
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

func day7() {
	data := ParseInputData("data/d7.txt")

	sumsPart2 := 0

	for index, v := range data {

	}

	fmt.Println("Day 6 Part 1: ", len(LocMap2))
	fmt.Println("Day 6 Part 2: ", sumsPart2)
}

func traversePath(data []string, start Loc, startPath Path, previousLoc []Loc) (bool, []Loc) {

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
	locs = append(locs, previousLoc...)
	traverse := true
	//first := 0

	for traverse {
		//startY, startX := findCurrentPos(data)
		if start.x == -1 || start.y == -1 {
			break
		}
		//dire := string(data[start.y][start.x])

		//rowToUpdate := data[start.y]
		//cpyRow := strings.Clone(rowToUpdate)
		//updatedRow := strings.Replace(rowToUpdate, dire, ".", 1)
		y, x, locs2 := findBlocker(start.x, start.y, startPath, data)
		startPath = directionMap[startPath.Next]
		start = Loc{x: x, y: y}
		//y, x, locs2 := findPath(directionMap, data)
		//fmt.Println("path locs:", locs)
		//fmt.Println("new locs2:", locs2)
		if len(locs2) > 1 {
			for c := range slices.Chunk(locs2, 2) {
				if slices.Contains(locs, c[0]) && slices.Contains(locs, c[1]) {
					// Contains inf loop
					/* 				fmt.Println("Found colliding path")
					fmt.Println(locs)
					fmt.Println(locs2)
					fmt.Println("") */

					return false, locs
				}
			}
			/* 			if slices.Contains(locs, locs2[0]) && slices.Contains(locs, locs2[1]) {
				// Contains inf loop
								fmt.Println("Found colliding path")
				   				fmt.Println(locs)
				   				fmt.Println(locs2)
				   				fmt.Println("")

				return false, locs
			} */
		}
		/* 		if slices.Contains(locs2, start) && first != 1 {
			// Contains inf loop
			//fmt.Println("Found start point")
			return false, locs
		} */
		/* 		for _, v := range locs2 {
			tmpData := slices.Clone(data)
			tmpUpdatedRow := strings.Replace(tmpData[startY], dire, ".", 1)
			tmpData[startY] = tmpUpdatedRow
			tmpData[v.y] = JoinUpdatedRow(tmpData[v.y], directionMap[dire].Next, v.x, v.y)
			fmt.Println(tmpData[v.y])
			a := traversePath(tmpData, v, []Loc{})
			if !a {
				return false
				//fmt.Println(tmpData[v.y])
				//fmt.Println("traverse result:", a, "For x:", v.x, "y:", v.y)
			}

			//fmt.Println()
		} */

		//fmt.Println("Start Loc", start.y, start.x)
		//fmt.Println(start.y, start.x, y, x, startPath)
		//fmt.Println(locs2)
		for _, v := range locs2 {
			if LocMap[v] == 0 {
				LocMap[v] = 1
			} else {
				LocMap[v] = LocMap[v] + 1
			}
			/* 			if LocMap[v] > 4171 {
				fmt.Println(LocMap[v], v)
				traverse = false
			} */
			/* 			if LocMap[v] > 1300 {
				//fmt.Println(LocMap[v])
				//fmt.Println(v)
				return false
			} */
		}
		// Part 2
		//sqLocs = append(sqLocs, Loc{x: x, y: y})
		/* 		if slices.Contains(locs2, start) && first != 1 {
		   			// Contains inf loop
		   			//fmt.Println("Found start point")
		   			return false, locs
		   		}

		   		if len(locs2) > 0 && slices.Contains(locs, locs2[0]) {
		   			//fmt.Println(LocMap)
		   			return false, locs
		   		} */
		/* 		existsInlocs := 0
		   		for _, v := range locs2 {
		   			if slices.Contains(locs, v) {
		   				existsInlocs++
		   			}
		   		}
		   		if existsInlocs == len(locs2) {
		   			return false
		   		} */

		locs = append(locs, locs2...)
		if y == -1 && x == -1 {
			traverse = false
			break
		}
		// remove old guard Char
		//data[start.y] = updatedRow
		// update new guard Char
		//data[y] = JoinUpdatedRow(data[y], directionMap[dire].Next, x, y)

		//fmt.Println("x:", x, "y:", y)
		//first = 1
	}
	//first = 1
	return true, locs
}

func JoinUpdatedRow(startRow, newDirection string, x, y int) string {
	splitString := strings.Split(startRow, "")
	startString := strings.Join(splitString[:x], "")
	startString = startString + newDirection
	return startString + strings.Join(splitString[x+1:], "")
}

func findPath(directions map[string]Path, data []string) (int, int, []Loc) {
	locs := make([]Loc, 0)
	y, x := findCurrentPos(data)
	dire := string(data[y][x])
	var maxDistance int
	if directions[dire].yDir != 0 {
		maxDistance = len(data)
	} else {
		maxDistance = len(data[0])
	}
	travelledDistance := 0
	for i := 1; i < maxDistance; i++ {
		yloc := y + (directions[dire].yDir * i)
		xloc := x + (directions[dire].xDir * i)
		if yloc < 0 || yloc >= len(data) {
			continue
		}
		if xloc < 0 || xloc >= len(data[0]) {
			continue
		}
		value := string(data[yloc][xloc])
		//fmt.Println(value)
		if value == "#" {
			return yloc + directions[dire].yDir*(-1), xloc + directions[dire].xDir*(-1), locs
		}
		locs = append(locs, Loc{x: xloc, y: yloc})
		travelledDistance++
	}
	return -1, -1, locs
}
func findBlocker(x, y int, direction Path, data []string) (int, int, []Loc) {
	locs := make([]Loc, 0)
	//y, x := findCurrentPos(data)
	var maxDistance int
	if direction.yDir != 0 {
		maxDistance = len(data)
	} else {
		maxDistance = len(data[0])
	}
	travelledDistance := 0
	for i := 1; i < maxDistance; i++ {
		yloc := y + (direction.yDir * i)
		xloc := x + (direction.xDir * i)
		if yloc < 0 || yloc >= len(data) {
			continue
		}
		if xloc < 0 || xloc >= len(data[0]) {
			continue
		}
		value := string(data[yloc][xloc])
		//fmt.Println(value)
		if value == "#" {
			return yloc + direction.yDir*(-1), xloc + direction.xDir*(-1), locs
		}
		locs = append(locs, Loc{x: xloc, y: yloc, direction: direction.Next})
		travelledDistance++
	}
	return -1, -1, locs
}

func findCurrentPos(data []string) (int, int) {
	for yIndex, v := range data {
		row := strings.Split(v, "")
		xIndex := slices.IndexFunc(row, func(n string) bool {
			return n == "<" || n == ">" || n == "v" || n == "^"
		})
		if xIndex != -1 {
			return yIndex, xIndex
		}
	}
	return -1, -1
}

func sortF(ints map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		if slices.Contains(ints[a], b) {
			return -1
		}
		return 0
	}
}

func isDataValid(ints []int) bool {
	order := make([]int, 0)
	for index, current := range ints {
		if index+1 >= len(ints) {
			order = append(order, current)
			toReverse := slices.Clone(order)
			slices.Reverse(toReverse)
			if slices.IsSorted(order) || slices.IsSorted(toReverse) {
				continue
			}
			return false
		}
		next := ints[index+1]
		isValid, _ := IsValidAndOrderedAsc(current, next)
		if !isValid {
			return false
		}

		order = append(order, current)
		toReverse := slices.Clone(order)
		slices.Reverse(toReverse)
		if slices.IsSorted(order) || slices.IsSorted(toReverse) {
			continue
		}
		return false
	}
	return true
}

func IsValidAndOrderedAsc(left int, right int) (bool, int) {
	if left-right > 3 || left-right < -3 || left-right == 0 {
		return false, -1
	}
	if left < right {
		return true, 1
	}
	return true, -1
}

func CountValues(list []float64, target float64) float64 {
	count := 0.0
	for _, v := range list {
		if v == target {
			count++
		}
	}
	return count
}

func ParseInputData(inputFile string) []string {
	data, error := os.ReadFile(inputFile)

	if error != nil {
		log.Fatal(error)
	}

	inputData := strings.Split(string(data[:]), "\r\n")
	return inputData
}
