package visuals

import (
	"fmt"
	"time"

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
	Foreground(lipgloss.Color("#9333ea")).
	MarginRight(1)

var partTimeStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4338ca"))

func PrintPart(part string, start time.Time, args ...any) {
	duration := time.Since(start)

	fmt.Print(partStyle.Underline(true).Render("Part " + part))
	fmt.Print(partTimeStyle.Render("(" + duration.String() + ")"))
	fmt.Print(partStyle.Render(":"))
	fmt.Println(args...)
}
