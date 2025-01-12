package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "doomsday-algorithm-cli",
}

func Execute() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(practiceCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
