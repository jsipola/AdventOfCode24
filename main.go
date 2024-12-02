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
	safeCountPart2 := 0
rowLoop:
	for _, v := range data {
		fields := strings.Fields(v)
		intFields := make([]int, len(fields))
		for i, v := range fields {
			intFields[i], _ = strconv.Atoi(v)
		}
		//dampener := 0
		isAscRow := "NONE"
		//fmt.Println()
		//countLoop:
		for index, current := range intFields {
			if index+1 >= len(intFields) {
				continue
			}
			//fmt.Println(intFields[0 : index+2])

			next := intFields[index+1]
			isValid, isAscPair := IsValidAndOrderedAsc(current, next)
			if isAscRow == "NONE" {
				isAscRow = isAscPair
			}
			if !isValid || isAscRow != isAscPair {
				/* 				if index-1 >= 0 {
				   					isValidLeftOfIndex, _ := IsValidAndOrderedAsc(intFields[index-1], intFields[index])
				   					if !isValidLeftOfIndex && dampener == 0 {
				   						dampener = 1
				   						intFields = append(intFields[:index], intFields[index+1:]...)
				   						fmt.Println("Bad level (isValidLeftOfIndex)", fields)
				   						fmt.Println("Bad level (isValidLeftOfIndex)", intFields)
				   						goto countLoop
				   					}
				   				}
				   				if index+1 <= len(intFields) {
				   					isValidRightOfIndex, _ := IsValidAndOrderedAsc(intFields[index], intFields[index+1])
				   					if !isValidRightOfIndex && dampener == 0 {
				   						dampener = 1
				   						intFields = append(intFields[:index+1], intFields[index+2:]...)
				   						fmt.Println("Bad level (isValidRightOfIndex)", fields)
				   						fmt.Println("Bad level (isValidRightOfIndex)", intFields)
				   						goto countLoop
				   					}
				   				} */

				continue rowLoop
			}
			// check if next value not either Min or Max so that the row is neither
			// ordered ASC or DESC
			/* 			reversed := slices.Clone(intFields[0 : index+2])
			   			slices.Reverse(reversed)
			   			if !slices.IsSorted(intFields[0:index+2]) && !slices.IsSorted(reversed) {
			   				if dampener != 0 {
			   					// Exit checking row values since more than 1 invalid levels
			   					fmt.Println("Bad level (order)", fields)
			   					fmt.Println("Bad level (order)", intFields)
			   					continue rowLoop
			   				}
			   				dampener = 1
			   				fmt.Println("removed (order) ", intFields[index+1])
			   				intFields = append(intFields[:index+1], intFields[index+2:]...)
			   				goto countLoop
			   			} */

			/* 			if v2-next > 3 || v2-next < -3 || v2-next == 0 {
			   				if dampener != 0 {
			   										fmt.Println()
			   					   					fmt.Println(v2)
			   					   					fmt.Println(next)
			   					fmt.Println("Bad level (large)", fields)
			   					fmt.Println("Bad level (large)", intFields)
			   					continue rowLoop
			   				}
			   				dampener = 1
			   				fmt.Println("removed (large) ", intFields[index])
			   				intFields = append(intFields[:index+1], intFields[index+2:]...)
			   				goto countLoop
			   			}
			   						fmt.Println("Next ", next)
			   			   			fmt.Println("index ", index)
			   			   			fmt.Println("Max ", next != slices.Max(intFields[0:index+2]))
			   			   			fmt.Println("Min ", next != slices.Min(intFields[0:index+2]))

			   			if next != slices.Max(intFields[0:index+2]) && next != slices.Min(intFields[0:index+2]) {
			   				if dampener != 0 {
			   					fmt.Println()
			   					   fmt.Println(v2)
			   					   fmt.Println(next)
			   					fmt.Println("Bad level (order)", fields)
			   					fmt.Println("Bad level (order)", intFields)
			   					continue rowLoop
			   				}
			   				dampener = 1
			   				fmt.Println("removed (order) ", intFields[index])
			   				intFields = append(intFields[:index], intFields[index+1:]...)
			   				goto countLoop
			   			} */
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
	fmt.Printf("Day 2 Part 2: %d\n", safeCountPart2)

}

func IsValidAndOrderedAsc(left int, right int) (bool, string) {
	if left-right > 3 || left-right < -3 || left-right == 0 {
		return false, "DESC"
	}
	if left < right {
		return true, "ASC"
	}
	return true, "DESC"
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
