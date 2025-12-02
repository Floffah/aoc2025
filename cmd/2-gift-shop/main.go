package main

import (
	"strconv"
	"strings"

	"github.com/Floffah/aoc2025/internal/pkg/inputs"
	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func main() {
	visuals.PrintDay("2")

	input := inputs.GetInput(2)
	ranges := strings.Split(strings.TrimSpace(input), ",")

	part1(ranges)
	part2(ranges)
}

func part1(ranges []string) {
	var sum int64 = 0

	for _, r := range ranges {
		start, end := parseRange(r)

		for i := start; i <= end; i++ {
			numStr := strconv.FormatInt(i, 10)

			lo := numStr[:len(numStr)/2]
			hi := numStr[len(numStr)/2:]

			if lo == hi {
				sum += i
			}
		}
	}

	visuals.PrintPart("1", "Sum of invalid gift shop ranges:", sum)
}

func part2(ranges []string) {
	var sum int64 = 0

	for _, r := range ranges {
		start, end := parseRange(r)

		for i := start; i <= end; i++ {
			numStr := strconv.FormatInt(i, 10)
			parts := splitEvenly(numStr)
			if parts != nil && allEqual(parts) && allEqualLen(parts) {
				sum += i
			}
		}
	}

	visuals.PrintPart("2", "Sum of invalid gift shop ranges with more than two equal parts:", sum)
}

func splitEvenly(s string) []string {
	for i := 2; i <= len(s); i += 1 {
		parts := splitIntoParts(s, i)
		if parts != nil && allEqual(parts) {
			return parts
		}
	}

	return nil
}

func splitIntoParts(s string, parts int) []string {
	if len(s)%parts != 0 {
		return nil
	}

	partLen := len(s) / parts
	result := make([]string, parts)

	for i := 0; i < parts; i++ {
		result[i] = s[i*partLen : (i+1)*partLen]
	}

	return result
}

func allEqual(s []string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}

	return true
}

func allEqualLen(s []string) bool {
	for i := 1; i < len(s); i++ {
		if len(s[i]) != len(s[0]) {
			return false
		}
	}

	return true
}

func parseRange(r string) (int64, int64) {
	bounds := strings.Split(r, "-")
	if len(bounds) != 2 {
		panic("invalid range")
	}

	start, err := strconv.ParseInt(strings.TrimSpace(bounds[0]), 10, 64)
	if err != nil {
		panic(err)
	}

	end, err := strconv.ParseInt(strings.TrimSpace(bounds[1]), 10, 64)
	if err != nil {
		panic(err)
	}

	return start, end
}
