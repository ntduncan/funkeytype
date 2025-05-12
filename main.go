package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	typetest "ntduncan.com/typer/type-test"
)

type Model struct {
	cursor int
	test   typetest.TypeTest
}

func InitModel() Model {
	tt := typetest.New("This is the type test")

	m := Model{
		cursor: 0,
		test:   tt,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "tab":
			//restart
		default:
			//Handle normal keypress
			if m.test.Wpm == 0 {
				m.test.HandleKeyPress(msg.String(), m.cursor)
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	white := lipgloss.Color("#FFF")
	s := strings.Builder{}

	title := lipgloss.NewStyle().Align(lipgloss.Center).Foreground(white).Render("\n\nSTART TYPING TO TYPE TEST\n\n\n")
	s.WriteString(title)

	body := ""

	correct := lipgloss.NewStyle().Foreground(lipgloss.Color("#85DEAD"))
	incorrect := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color("#CE2029"))
	cursor := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color("#1E90FF"))
	blank := lipgloss.NewStyle().Foreground(lipgloss.Color("#a7a7a"))

	for i, p := range m.test.Params {
		if i == m.cursor {
			body += cursor.Render(p.Char)
		} else if p.IsValid {
			body += correct.Render(p.Char)
		} else if !p.IsValid && p.Input != "" {
			body += incorrect.Render(p.Char)
		} else {
			body += blank.Render(p.Char)
		}
	}

	s.WriteString(body)

	return s.String()
}

func main() {
	p := tea.NewProgram(InitModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Exited with error: %s", err)
		os.Exit(1)
	}
	fmt.Println("Silly boy. You forgot to implement your main function")
}
