package cmd

import (
	"fmt"
	"time"

	"github.com/ArturMinelli/doomsday-algorithm-cli/doomsday"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run YYYY-MM-DD",
	Short: "Calculate the weekday for a given date using the Doomsday algorithm",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		date, err := time.Parse("2006-01-02", args[0])
		if err != nil {
			fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
			return
		}

		result := doomsday.Run(date)

		fmt.Println(result)
	},
}
