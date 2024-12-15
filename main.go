package adventofcode24

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Loc struct {
	x         int
	y         int
	direction string
}

type Path struct {
	yDir int
	xDir int
	Next string
}

func convertToInts(s string) []int {
	strs := strings.Split(s, "")
	ints := make([]int, 0)
	for _, v := range strs {
		val, _ := strconv.Atoi(v)
		ints = append(ints, val)
	}
	return ints
}

func moveblocks(block []string, final *[]string) {
	emptySlot := FindEmptyFunc(block)
	if emptySlot == -1 {
		*final = append(*final, block...)
		return
	}
	if emptySlot > 0 {
		*final = append(*final, block[0:emptySlot]...)
		moveblocks(block[emptySlot:], final)
		return
	} else {
		for a, b := range slices.Backward(block) {
			if b != "." {
				*final = append(*final, block[a])
				moveblocks(block[emptySlot+1:a], final)
				break
			}
		}
	}
}

func FindEmptyFunc(strs []string) int {
	return slices.IndexFunc(strs, func(str string) bool {
		return str == "."
	})
}
func FindNotEmptyFunc(strs []string) int {
	return slices.IndexFunc(strs, func(str string) bool {
		return str != "."
	})
}

func isValid(targetSum int, nums []int, isPart2 bool) bool {
	if len(nums) == 1 {
		return nums[0] == targetSum
	} else if nums[0] > targetSum {
		return false
	}

	if isValid(targetSum, append([]int{nums[0] + nums[1]}, nums[2:]...), isPart2) {
		return true
	}
	if isValid(targetSum, append([]int{nums[0] * nums[1]}, nums[2:]...), isPart2) {
		return true
	}
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[1]))
	return isValid(targetSum, append([]int{concat}, nums[2:]...), isPart2) && isPart2
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
