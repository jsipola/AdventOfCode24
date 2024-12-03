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
