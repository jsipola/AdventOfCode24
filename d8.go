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

func deleteOoBAntenna(current []antenna, height, length int) []antenna {
	return slices.DeleteFunc(current, func(item antenna) bool {
		if item.y >= height || item.y < 0 {
			return true
		}
		if item.x >= length || item.x < 0 {
			return true
		}
		return false
	})
}

type antenna struct {
	x int
	y int
}

func findAllAntennas(data []string, marker rune) []antenna {
	allAntennas := make([]antenna, 0)
	for yIndex, v := range data {
		antennas := findAntenna(v, marker)
		if len(antennas) > 0 {
			for _, v := range antennas {
				allAntennas = append(allAntennas, antenna{y: yIndex, x: v})
			}

		}
	}

	antiNodes := make([]antenna, 0)
	for index, v := range allAntennas {
		a := slices.Clone(allAntennas)
		getAllAntiNodes(v, append(a[:index], a[index+1:]...), &antiNodes)
	}
	/* fmt.Println("AntiNodes: ", antiNodes) */
	return antiNodes
}

func getAllAntiNodes(start antenna, rest []antenna, antiNodes *[]antenna) {
	if len(rest) == 0 {
		return
	}
	if len(rest) == 2 {
		*antiNodes = append(*antiNodes, start)
	}
	antiNode := calculateAntiNode(start, rest[0])
	*antiNodes = append(*antiNodes, antiNode)
	// TODO parse width length dynamically
	if start.x <= 50 && start.y <= 50 && start.y >= 0 && start.x >= 0 {
		getAllAntiNodes(rest[0], []antenna{antiNode}, antiNodes)
	}
	getAllAntiNodes(start, rest[1:], antiNodes)
}

func calculateAntiNode(a, b antenna) antenna {
	x := (b.x - a.x) + b.x
	y := (b.y - a.y) + b.y
	return antenna{x: x, y: y}
}

func findAntenna(data string, marker rune) []int {
	antennas := make([]int, 0)
	/* 	xIndex := strings.IndexFunc(data, func(c rune) bool {
	   		return c == marker
	   	})
	   	return xIndex */
	for xIndex, v := range data {
		if v == marker {
			antennas = append(antennas, xIndex)
		}
	}
	return antennas
}
