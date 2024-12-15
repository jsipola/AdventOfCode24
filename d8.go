package adventofcode24

import (
	"fmt"
	"slices"
	"unicode"
)

func day8() {
	data := ParseInputData("data/d8.txt")

	f := func(s rune) bool {
		if unicode.IsDigit(s) || unicode.IsLetter(s) {
			return true
		}
		return false
	}
	allsums := 0
	exclude := make([]rune, 0)
	allAntennas := make([]antenna, 0)
	for _, v := range data {
		for _, char := range v {
			if f(char) && !slices.Contains(exclude, char) {
				allAntennas = append(allAntennas, findAllAntennas(data, char)...)
				exclude = append(exclude, char)
			}
		}
	}

	uniqueAntenna := make(map[antenna]int, 0)
	filtered := deleteOoBAntenna(allAntennas, len(data), len(data[0]))
	for _, v := range filtered {
		uniqueAntenna[v] = 1
	}

	fmt.Println("Day 8 Part 1:    ", len(uniqueAntenna))
	fmt.Println("Day 8 Part 2:    ", allsums)
}
