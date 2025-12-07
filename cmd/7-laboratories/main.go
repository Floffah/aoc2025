package main

import (
	"strings"
	"time"

	"github.com/Floffah/aoc2025/internal/pkg/inputs"
	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func main() {
	visuals.PrintDay("7")

	input := inputs.GetInput(7)
	lines := strings.Split(input, "\n")

	linesPart1 := make([]string, len(lines))
	copy(linesPart1, lines)
	linesPart2 := make([]string, len(lines))
	copy(linesPart2, lines)

	part1(linesPart1)
	part2(linesPart2)
}

func part1(lines []string) {
	start := time.Now()

	timesSplit := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'S' {
				lines[y] = lines[y][:x] + "|" + lines[y][x+1:]
			}

			if y > 0 {
				if lines[y-1][x] == '|' {
					if lines[y][x] == '^' {
						timesSplit++

						lines[y] = lines[y][:x-1] + "|^|" + lines[y][x+2:]
					} else if lines[y][x] == '.' {
						lines[y] = lines[y][:x] + "|" + lines[y][x+1:]
					}
				}
			}
		}
	}

	visuals.PrintPart("1", start, "Times beam split:", timesSplit)
}

func part2(lines []string) {
	start := time.Now()

	height := len(lines)
	width := len(lines[0])

	sx, sy := -1, -1
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if lines[y][x] == 'S' {
				sx, sy = x, y
				break
			}
		}
		if sx != -1 {
			break
		}
	}

	count := make([]int64, width)
	next := make([]int64, width)

	count[sx] = 1

	for y := sy + 1; y < height; y++ {
		for i := range next {
			next[i] = 0
		}

		for x := 0; x < width; x++ {
			c := count[x]
			if c == 0 {
				continue
			}

			cell := lines[y][x]
			if cell == '^' {
				next[x-1] += c
				next[x+1] += c
			} else {
				next[x] += c
			}
		}

		count, next = next, count
	}

	var timelines int64
	for _, c := range count {
		timelines += c
	}

	visuals.PrintPart("2", start, "Timelines:", timelines, "(95408386769474)")
}
