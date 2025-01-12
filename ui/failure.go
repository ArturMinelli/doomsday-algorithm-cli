package ui

import (
	"fmt"
	"time"

	"github.com/ArturMinelli/doomsday-algorithm-cli/doomsday"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FailureModel struct {
	Doomsday doomsday.Doomsday
	Guess    int
	Elapsed  time.Duration
	width    int
	height   int
}

func NewFailure(doomsday doomsday.Doomsday, guess int, elapsed time.Duration) FailureModel {
	return FailureModel{
		Doomsday: doomsday,
		Guess:    guess,
		Elapsed:  elapsed,
	}
}

func (m FailureModel) Init() tea.Cmd {
	return nil
}

func (m FailureModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m FailureModel) View() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF4136")).
		Render("Incorrect!")

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
			{"", ""},
			{"Century", fmt.Sprintf("%d", m.Doomsday.Variables.Century)},
			{"Decade", fmt.Sprintf("%d", m.Doomsday.Variables.Decade)},
			{"Month Doomsday", fmt.Sprintf("%d", m.Doomsday.Variables.MonthDoomsday)},
			{"", ""},
			{"Day to Month Doomsday Offset", fmt.Sprintf("%d", m.Doomsday.Variables.DayToMonthDoomsdayOffset)},
			{"Division Decade by Twelve", fmt.Sprintf("%d", m.Doomsday.Variables.DivisionDecadeByTwelve)},
			{"Remainder Decade by Twelve", fmt.Sprintf("%d", m.Doomsday.Variables.RemainderDecadeByTwelve)},
			{"Division Remainder by Four", fmt.Sprintf("%d", m.Doomsday.Variables.DivisionRemainderByFour)},
			{"Century Code", fmt.Sprintf("%d", m.Doomsday.Variables.CenturyCode)},
		}),
	)
	variablesTableStyle := table.DefaultStyles()
	variablesTableStyle.Cell = variablesTableStyle.Cell.Align(lipgloss.Center).Margin(0, 0).Padding(0, 0).Foreground(lipgloss.Color("#FF9999"))
	variablesTableStyle.Selected = variablesTableStyle.Selected.Foreground(lipgloss.Color("#FF9999"))
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
