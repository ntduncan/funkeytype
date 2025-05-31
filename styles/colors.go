package styles

import "github.com/charmbracelet/lipgloss"

type ColorsType struct {
	White  lipgloss.Color
	Black  lipgloss.Color
	Orange lipgloss.Color
	Red    lipgloss.Color
	Blue   lipgloss.Color
}

var Colors = ColorsType{
	White:  lipgloss.Color("#FFF"),
	Black:  lipgloss.Color("#111"),
	Orange: lipgloss.Color("#CD5C2A"),
	Red:    lipgloss.Color("#CE2029"),
	Blue:   lipgloss.Color("#0018F9"),
}
