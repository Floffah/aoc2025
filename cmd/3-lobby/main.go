package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/Floffah/aoc2025/internal/pkg/inputs"
	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func main() {
	visuals.PrintDay("3")

	input := inputs.GetInput(3)
	lines := strings.Split(input, "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	start := time.Now()

	sum := 0

	for _, line := range lines {
		numStr := strings.TrimSpace(line)
		highest := highestPossibleJoltageGreedy(numStr, 2)
		sum += highest
	}

	visuals.PrintPart("1", start, "Sum of highest possible joltages (2 combins):", sum)
}

func part2(lines []string) {
	start := time.Now()

	sum := 0

	for _, line := range lines {
		numStr := strings.TrimSpace(line)
		highest := highestPossibleJoltageGreedy(numStr, 12)
		sum += highest
	}

	visuals.PrintPart("2", start, "Sum of highest possible joltages (12 combins):", sum)
}

func highestPossibleJoltageGreedy(numStr string, combinsCount int) int {
	result := 0
	start := 0

	for pos := 0; pos < combinsCount; pos++ {
		end := len(numStr) - (combinsCount - pos)
		maxDigit := byte('0')
		maxIdx := start

		for i := start; i <= end; i++ {
			currentCol := numStr[i]
			if currentCol > maxDigit {
				maxDigit = currentCol
				maxIdx = i
				if maxDigit == '9' {
					break
				}
			}
		}

		result = result*10 + int(maxDigit-'0')
		start = maxIdx + 1
	}

	return result
}

func highestPossibleJoltageUnoptimised(numStr string, combinsCount int) int {
	bankHighest := 0

	getNumTraverseFrom(numStr, "", 0, 0, combinsCount, &bankHighest)

	return bankHighest
}

func getNumTraverseFrom(numStr, parentCombinStr string, index, depth, combinsCount int, bankHighest *int) {
	for i := index; i < len(numStr); i++ {
		if depth >= combinsCount {
			return
		}

		units := string(numStr[i])
		combinedStr := parentCombinStr + units

		if depth == combinsCount-1 {
			combined, err := strconv.Atoi(combinedStr)
			if err != nil {
				panic(err)
			}

			if combined > *bankHighest {
				*bankHighest = combined
			}
		}

		getNumTraverseFrom(numStr, combinedStr, i+1, depth+1, combinsCount, bankHighest)
	}
}

func firstVer_highestPossible2dJoltage(numStr string) int {
	bankHighest := 0

	for i := 0; i < len(numStr); i++ {
		tens := string(numStr[i])

		for j := i + 1; j < len(numStr); j++ {

			units := string(numStr[j])
			combinedStr := tens + units
			combined, err := strconv.Atoi(combinedStr)
			if err != nil {
				panic(err)
			}

			if combined > bankHighest {
				bankHighest = combined
			}
		}
	}

	return bankHighest
}
