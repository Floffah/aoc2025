package main

import (
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/Floffah/aoc2025/internal/pkg/inputs"
	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func main() {
	visuals.PrintDay("5")

	input := inputs.GetInput(5)
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rangesLines := strings.Split(sections[0], "\n")
	idsLines := strings.Split(sections[1], "\n")

	part1(rangesLines, idsLines)
	part2(rangesLines)
}

func part1(rangesLines, idsLines []string) {
	start := time.Now()
	ranges := parseRanges(rangesLines)

	sum := 0

	for _, idLine := range idsLines {
		id, _ := strconv.Atoi(strings.TrimSpace(idLine))

		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				sum += 1
				break
			}
		}
	}

	visuals.PrintPart("1", start, "Fresh IDs count:", sum)
}

func part2(rangesLines []string) {
	start := time.Now()
	ranges := parseRanges(rangesLines)

	sum := 0

	rangesCmp := func(a, b []int) int {
		return a[0] - b[0]
	}
	slices.SortFunc(ranges, rangesCmp)

	mergedRanges := make([][]int, 0, len(ranges))

	for _, r := range ranges {
		if len(mergedRanges) == 0 {
			mergedRanges = append(mergedRanges, r)
			continue
		}

		lastRange := mergedRanges[len(mergedRanges)-1]

		if r[0] <= lastRange[1]+1 {
			if r[1] > lastRange[1] {
				lastRange[1] = r[1]
			}
		} else {
			mergedRanges = append(mergedRanges, r)
		}
	}

	for _, r := range mergedRanges {
		sum += r[1] - r[0] + 1
	}

	visuals.PrintPart("2", start, "Total fresh IDs in ranges:", sum)
}

func parseRanges(rangesLines []string) [][]int {
	result := make([][]int, 0, len(rangesLines))

	for _, rline := range rangesLines {
		parts := strings.Split(rline, "-")
		if len(parts) != 2 {
			continue
		}

		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		result = append(result, []int{start, end})
	}

	return result
}
