package main

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/Floffah/aoc2025/internal/pkg/inputs"
	"github.com/Floffah/aoc2025/internal/pkg/visuals"
)

func main() {
	visuals.PrintDay("1")

	input := inputs.GetInput(1)
	lines := strings.Split(input, "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	start := time.Now()

	bearing := 50
	sum := 0

	for _, line := range lines {
		parseAndComputeBearing(line, &bearing)

		if (bearing % 100) == 0 {
			sum += 1
		}
	}
	visuals.PrintPart("1", start, "0 occurs", sum, "times")
}

//       __
//     _|==|_
//      ('')___/
//  >--(`^^')
//    (`^'^'`)
//    `======'

func part2(lines []string) {
	start := time.Now()

	bearing := 50
	sum := 0

	for _, line := range lines {
		previousBearing := bearing

		parseAndComputeBearing(line, &bearing)

		var diff float64

		if bearing > previousBearing {
			diff = math.Floor(float64(bearing)/100) - math.Floor(float64(previousBearing)/100)
		} else if bearing < previousBearing {
			diff = math.Floor(float64(previousBearing-1)/100) - math.Floor(float64(bearing-1)/100)
		} else {
			diff = 0
		}

		sum += int(diff)
	}

	visuals.PrintPart("2", start, "0 is hit", sum, "times")
}

func parseAndComputeBearing(line string, bearing *int) {
	prefixChar := line[0]
	num, err := strconv.ParseInt(line[1:], 10, 32)
	if err != nil {
		panic(err)
	}

	switch prefixChar {
	case 'L':
		*bearing -= int(num)
	case 'R':
		*bearing += int(num)
	default:
		panic("unknown prefix")
	}
}
