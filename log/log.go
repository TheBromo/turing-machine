package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const width = 100

var styleNormal = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#37d99e")).
	Background(lipgloss.Color("#101317")).
	PaddingTop(1).
	PaddingBottom(1).
	PaddingLeft(4).
	PaddingRight(4).
	Width(width).BorderStyle(lipgloss.RoundedBorder())

var styleError = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#e87979")).
	Background(lipgloss.Color("#101317")).
	Width(width)

func Print(text string) {
	fmt.Println(styleNormal.Render(text))
}

func PrintError(text string) {
	fmt.Println(styleError.Render(text))
}
