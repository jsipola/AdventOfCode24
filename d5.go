package adventofcode24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

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
