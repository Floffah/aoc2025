package inputs

import (
	"fmt"
	"os"
)

func GetInput(day int, args ...int) string {
	fileName := fmt.Sprint("assets/", day, ".input")
	if len(args) > 0 {
		fileName = fmt.Sprint("assets/", day, "-", args[0], ".input")
	}

	inputFile, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(inputFile)
}
