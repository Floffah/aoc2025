package visuals

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var dayStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#ffffff")).
	Background(lipgloss.Color("#7c3aed")).
	Width(20).
	Align(lipgloss.Center).
	MarginTop(1)

func PrintDay(day string) {
	fmt.Println(dayStyle.Render("Day " + day))
}

var partStyle = lipgloss.NewStyle().
	Underline(true).
	Foreground(lipgloss.Color("#9333ea")).
	MarginRight(1)

func PrintPart(part string, args ...any) {
	fmt.Print(partStyle.Render("Part " + part + ":"))
	fmt.Println(args...)
}
