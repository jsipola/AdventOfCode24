package adventofcode24

import (
	"fmt"
	"strconv"
	"strings"
)

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
