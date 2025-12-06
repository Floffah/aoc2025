package main

import (
	"strings"

	"github.com/Floffah/aoc2025/internal/pkg/inputs"
	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func main() {
	visuals.PrintDay("6")

	input := inputs.GetInput(6)
	lines := strings.Split(input, "\n")

	part1(lines)
	part2(lines)
}
