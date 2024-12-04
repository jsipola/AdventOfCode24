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

	inputData := strings.Split(string(data[:]), "\n")
	return inputData
}
