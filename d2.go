package adventofcode24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

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
