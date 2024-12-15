package adventofcode24

import (
	"fmt"
	"strconv"
	"strings"
)

func day7() {
	data := ParseInputData("data/d7.txt")

	type row struct {
		target int
		nums   []int
	}
	rows := make([]row, len(data))

	for index, v := range data {
		strs := strings.Split(v, ":")
		sum, _ := strconv.Atoi(strs[0])
		strs = strings.Fields(strs[1])
		ints := make([]int, len(strs))
		for index, v1 := range strs {
			ints[index], _ = strconv.Atoi(string(v1))
		}
		rows[index] = row{target: sum, nums: ints}
	}

	sums := 0
	for _, row := range rows {
		if isValid(row.target, row.nums, false) {
			sums += row.target
		}
	}

	allsums := 0
	for _, row := range rows {
		if isValid(row.target, row.nums, true) {
			allsums += row.target
		}
	}

	fmt.Println("Day 7 Part 1:    ", sums)
	fmt.Println("Day 7 Part 2:    ", allsums)
}
