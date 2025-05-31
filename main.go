package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"ntduncan.com/typer/styles"
	typetest "ntduncan.com/typer/type-test"
)

type Model struct {
	cursor   int
	viewport viewport.Model
	test     typetest.TypeTest
}

var BestWPM string = ""

func InitModel(width int, height int, size int) Model {
	tt := typetest.New(size)

	m := Model{
		cursor:   0,
		test:     tt,
		viewport: viewport.New(width, height),
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
			m = InitModel(m.viewport.Width, m.viewport.Height, m.test.Size)
		case "+", "plus":
			var newSize int

			switch m.test.Size {
			case 10:
				newSize = 25
			case 25:
				newSize = 50
			case 50:
				newSize = 100
			default:
				newSize = 10
			}
			m = InitModel(m.viewport.Width, m.viewport.Height, newSize)
		case "backspace":
			if m.cursor > 0 && m.cursor < len(m.test.Params) {
				m.test.Params[m.cursor].Input = ""
				m.test.Params[m.cursor].IsValid = false
				m.cursor--
			}

		default:
			//Handle normal keypress
			if m.test.EndTime.IsZero() && m.cursor != len(m.test.Params) {
				m.test.HandleKeyPress(msg.String(), m.cursor)
				if m.cursor != len(m.test.Params)-1 {
					m.cursor++
				}
			}
		}
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height
		m.viewport.Width = msg.Width
	}

	return m, nil
}

func (m Model) View() string {
	colors := styles.Colors

	title := lipgloss.NewStyle().
		Foreground(colors.Orange).
		Bold(true).
		BorderRight(true).
		BorderStyle(lipgloss.DoubleBorder()).
		Padding(0, 1, 0, 0).
		Margin(0, 1).
		Render("FunKeyType")

	wpm := m.test.GetWPM()
	wpmStyled := lipgloss.
		NewStyle().
		Width(m.viewport.Width).
		Align(lipgloss.Center).
		Render("WPM: " + wpm)

	if wpm > BestWPM {
		BestWPM = wpm
	}

	bestWPMStyled := lipgloss.
		NewStyle().
		Foreground(colors.Orange).
		Render(BestWPM)

	testLen := lipgloss.NewStyle().
		Bold(true).
		BorderRight(true).
		BorderStyle(lipgloss.DoubleBorder()).
		Padding(0, 1, 0, 0).
		Margin(0, 1).
		Render(m.test.GetTestSize())

	topBar := lipgloss.
		NewStyle().
		Bold(true).
		Width(m.viewport.Width-2).
		Border(lipgloss.DoubleBorder()).
		Padding(0, 1).
		Render(title + testLen + "Top Score: " + bestWPMStyled)

	body := ""

	correct := lipgloss.NewStyle().Foreground(lipgloss.Color("#85DEAD"))
	incorrect := lipgloss.NewStyle().Foreground(colors.White).Background(lipgloss.Color(colors.Red))
	blockCursor := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF")).Background(colors.Black)
	lineCursor := lipgloss.NewStyle().Underline(true)
	blank := lipgloss.NewStyle()

	for i, p := range m.test.Params {
		if i == m.cursor {
			if (m.test.EndTime != time.Time{}) {
				body += blockCursor.Render(p.Char)
			} else {
				body += lineCursor.Render(p.Char)
			}
			continue
		} else if p.IsValid {
			body += correct.Render(p.Char)
			continue
		} else if !p.IsValid && p.Input != "" {
			body += incorrect.Render(p.Char)
			continue
		} else {
			body += blank.Render(p.Char)
			continue
		}
	}

	body = lipgloss.
		NewStyle().
		Height(m.viewport.Height-10).
		Width(m.viewport.Width-2).
		Align(lipgloss.Left).
		Padding(0, 10).
		Render(body)

	f := wordwrap.String(body, m.viewport.Width-10)
	m.viewport.SetContent(fmt.Sprintf("%v\n%v\n\n%v\n\n\n%v", topBar, wpmStyled, f, m.footer()))

	return m.viewport.View()
}

func (m Model) footer() string {
	f := strings.Builder{}

	cmdMenu := lipgloss.
		NewStyle().
		Width(m.viewport.Width - 2).
		Border(lipgloss.DoubleBorder()).
		Foreground(lipgloss.Color("#A7A7A7")).
		Render("\"esc\": Exit | \"tab\": New Test | \"+\": Test Length |")

	f.WriteString(cmdMenu)

	return f.String()
}

func main() {
	p := tea.NewProgram(InitModel(10, 10, 10), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Exited with error: %s", err)
		os.Exit(1)
	}
}
