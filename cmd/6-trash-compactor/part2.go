package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func part2(diffLenLines []string) {
	start := time.Now()

	lines := ensureAllLinesSameLength(diffLenLines)

	sum := 0

	currentOperator := ""
	var problemAccum []int
	//numAccum := ""

	for i := 0; i < len(lines[0]); i++ {
		verticalStr := strings.TrimSpace(getVerticalString(lines, i))

		if verticalStr == "" {
			if len(problemAccum) > 0 && currentOperator != "" {
				result := problemAccum[len(problemAccum)-1]

				for j := len(problemAccum) - 2; j >= 0; j-- {
					switch currentOperator {
					case "*":
						result = result * problemAccum[j]
					case "+":
						result = result + problemAccum[j]
					}
				}

				sum += result
				problemAccum = []int{}
				currentOperator = ""
			}
			continue
		}

		lastChar := verticalStr[len(verticalStr)-1]

		if lastChar == '*' || lastChar == '+' {
			currentOperator = string(lastChar)
			verticalStr = strings.TrimSpace(verticalStr[:len(verticalStr)-1])
		}

		num, err := strconv.Atoi(strings.TrimSpace(verticalStr))
		if err != nil {
			panic(err)
		}
		problemAccum = append(problemAccum, num)
	}

	visuals.PrintPart("2", start, "Sum of column results:", sum, "(10188206723429)")
}

func getVerticalString(lines []string, colIdx int) string {
	result := ""

	for _, line := range lines {
		if colIdx < len(line) {
			result += string(line[colIdx])
		}
	}

	return result
}

func ensureAllLinesSameLength(lines []string) []string {
	maxLen := 0

	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line) + 1
		}
	}

	result := make([]string, 0, len(lines))

	for _, line := range lines {
		paddedLine := fmt.Sprintf("%-*s", maxLen, line)
		result = append(result, paddedLine)
	}

	return result
}
