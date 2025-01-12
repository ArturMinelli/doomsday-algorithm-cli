package ui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ArturMinelli/doomsday-algorithm-cli/doomsday"
	"github.com/ArturMinelli/doomsday-algorithm-cli/speech"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TimerModel struct {
	Date      time.Time
	Guess     int
	stopwatch stopwatch.Model
	width     int
	height    int
}

func NewTimer() TimerModel {
	m := TimerModel{
		stopwatch: stopwatch.NewWithInterval(time.Millisecond * 100),
	}

	m.setGuess(-1)
	m.setDate(doomsday.NewRandomDate())

	return m
}

func (m TimerModel) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m TimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.setGuess(-1)
			return m, tea.Quit
		case "n":
			m.setGuess(-1)
			m.setDate(doomsday.NewRandomDate())

			return m, m.stopwatch.Reset()
		case "r":
			m.setGuess(-1)
			return m, m.stopwatch.Reset()
		case "0", "1", "2", "3", "4", "5", "6":
			guess, _ := strconv.Atoi(msg.String())
			m.setGuess(guess)
			return m, tea.Quit
		}

	default:
		var cmd tea.Cmd
		m.stopwatch, cmd = m.stopwatch.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m TimerModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFD700")).
		MarginBottom(1).
		Align(lipgloss.Center)

	timerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFD700")).
		MarginBottom(1).
		Align(lipgloss.Center)

	dateStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		MarginTop(1)

	optionsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#3498DB")).
		MarginTop(1)

	actionsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#95A5A6"))

	elapsed := strconv.FormatFloat(m.stopwatch.Elapsed().Seconds(), 'f', 1, 64)

	content := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s\n\n%s",
		titleStyle.Render("🎯 Practice Mode 🎯"),
		timerStyle.Render(elapsed),
		dateStyle.Render(m.Date.Format("2006-01-02")),
		optionsStyle.Render("0: Monday • 1: Tuesday • 2: Wednesday • 3: Thursday • 4: Friday • 5: Saturday • 6: Sunday"),
		actionsStyle.Render("Press 'q' to quit, 'n' for new date, 'r' to reset timer"),
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (m *TimerModel) setGuess(guess int) {
	m.Guess = guess
}

func (m *TimerModel) setDate(date time.Time) {
	go speech.Speak(date.Format("2006-01-02"))
	m.Date = date
}
