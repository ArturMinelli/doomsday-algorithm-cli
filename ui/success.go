package ui

import (
	"fmt"
	"time"

	"github.com/ArturMinelli/doomsday-algorithm-cli/doomsday"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SuccessModel struct {
	Doomsday doomsday.Doomsday
	Guess    int
	Elapsed  time.Duration
	width    int
	height   int
}

func NewSuccess(doomsday doomsday.Doomsday, guess int, elapsed time.Duration) SuccessModel {
	return SuccessModel{
		Doomsday: doomsday,
		Guess:    guess,
		Elapsed:  elapsed,
	}
}

func (m SuccessModel) Init() tea.Cmd {
	return nil
}

func (m SuccessModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "n", "q", " ", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m SuccessModel) View() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")).
		Render("Correct!")

	variablesTable := table.New(
		table.WithColumns([]table.Column{
			{Title: "", Width: 48},
			{Title: "", Width: 16},
		}),
		table.WithRows([]table.Row{
			{"Date", m.Doomsday.Date.Format("2006-01-02")},
			{"Correct Weekday", m.Doomsday.Date.Weekday().String()},
			{"Your Guess", fmt.Sprintf("%d", m.Guess)},
			{"Elapsed Time", m.Elapsed.String()},
		}),
	)
	variablesTableStyle := table.DefaultStyles()
	variablesTableStyle.Cell = variablesTableStyle.Cell.Align(lipgloss.Center).Margin(0, 0).Padding(0, 0).Foreground(lipgloss.Color("#99FF99"))
	variablesTableStyle.Selected = variablesTableStyle.Selected.Foreground(lipgloss.Color("#99FF99"))
	variablesTable.SetStyles(variablesTableStyle)

	actions := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#95A5A6")).
		Render("Press 'n', 'q' or space to continue")

	content := fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		title,
		variablesTable.View(),
		actions,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
