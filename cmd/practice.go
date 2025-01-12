package cmd

import (
	"github.com/ArturMinelli/doomsday-algorithm-cli/doomsday"
	"github.com/ArturMinelli/doomsday-algorithm-cli/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var practiceCmd = &cobra.Command{
	Use:   "practice",
	Short: "Practice the Doomsday algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			t := tea.NewProgram(
				ui.NewTimer(),
				tea.WithAltScreen(),
				tea.WithMouseCellMotion(),
			)

			result, err := t.Run()
			if err != nil {
				panic(err)
			}

			state := result.(ui.TimerModel)

			if state.Guess == -1 {
				break
			}

			doom := doomsday.Run(state.Date)

			if state.Guess == doom.Weekday {
				s := tea.NewProgram(
					ui.NewSuccess(doom, state.Guess),
					tea.WithAltScreen(),
					tea.WithMouseCellMotion(),
				)

				_, err := s.Run()
				if err != nil {
					panic(err)
				}
			} else {
				f := tea.NewProgram(
					ui.NewFailure(doom, state.Guess),
					tea.WithAltScreen(),
					tea.WithMouseCellMotion(),
				)

				_, err := f.Run()
				if err != nil {
					panic(err)
				}
			}
		}
	},
}
