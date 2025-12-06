package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func part1(lines []string) {
	start := time.Now()

	parsedLines := parseLines(lines)

	sum := 0

	for i := 0; i < len(parsedLines[0]); i++ {
		column := extractWholeColumn(parsedLines, i)

		result, err := strconv.Atoi(column[0])
		if err != nil {
			panic(err)
		}
		operator := column[len(column)-1]

		for _, val := range column[1 : len(column)-1] {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			switch operator {
			case "*":
				result = result * num
			case "+":
				result = result + num
			}
		}

		sum += result
	}

	visuals.PrintPart("1", start, "Sum of column results:", sum)
}

func parseLines(lines []string) [][]string {
	result := make([][]string, 0, len(lines))

	for _, line := range lines {
		parsedLine := parseLine(line)
		result = append(result, parsedLine)
	}

	return result
}

func parseLine(lines string) []string {
	result := make([]string, 0)

	accumulator := ""
	for _, char := range strings.TrimSpace(lines) {
		if string(char) == " " {
			if accumulator != "" {
				result = append(result, accumulator)
				accumulator = ""
			}

			continue
		}

		accumulator += string(char)
	}

	if accumulator != "" {
		result = append(result, accumulator)
	}

	return result
}

func extractWholeColumn(lines [][]string, colIdx int) []string {
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		if colIdx < len(line) {
			result = append(result, line[colIdx])
		}
	}

	return result
}
