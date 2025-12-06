package main

import (
	"strings"
	"time"

	"github.com/Floffah/aoc2025/internal/pkg/inputs"
	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func main() {
	visuals.PrintDay("4")

	input := inputs.GetInput(4)
	lines := strings.Split(input, "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	start := time.Now()

	sum := 0

	for lidx, line := range lines {
		for cidx, char := range strings.TrimSpace(line) {
			if string(char) == "@" {
				if accessiblePaper(&lines, cidx, lidx) {
					sum++
				}
			}
		}
	}

	visuals.PrintPart("1", start, "Safe forklift access positions:", sum)
}

func part2(lines []string) {
	start := time.Now()

	sum := 0
	changesCount := 0

	for ok := true; ok; ok = changesCount > 0 {
		changesCount = 0

		for lidx, line := range lines {
			for cidx, char := range strings.TrimSpace(line) {
				if string(char) == "@" {
					if accessiblePaper(&lines, cidx, lidx) {
						lines[lidx] = lines[lidx][:cidx] + "." + lines[lidx][cidx+1:]
						sum++
						changesCount++
					}
				}
			}
		}
	}

	visuals.PrintPart("2", start, "Safe forklift access positions with removes:", sum)
}

func accessiblePaper(lines *[]string, cidx, lidx int) bool {
	adjacentCount := 0
	matrix := [][]string{
		{access(lines, cidx-1, lidx-1), access(lines, cidx, lidx-1), access(lines, cidx+1, lidx-1)},
		{access(lines, cidx-1, lidx), "", access(lines, cidx+1, lidx)},
		{access(lines, cidx-1, lidx+1), access(lines, cidx, lidx+1), access(lines, cidx+1, lidx+1)},
	}

	for _, adjacent := range matrix {
		for _, cell := range adjacent {
			if cell == "@" {
				adjacentCount++
			}
		}
	}

	return adjacentCount < 4
}

func access(lines *[]string, x, y int) string {
	if y < 0 || y >= len(*lines) {
		return ""
	}
	line := strings.TrimSpace((*lines)[y])
	if x < 0 || x >= len(line) {
		return ""
	}
	return string(line[x])
}
